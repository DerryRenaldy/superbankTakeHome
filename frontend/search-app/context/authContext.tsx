"use client"

import { createContext, useContext, useState, ReactNode, useEffect } from "react";
import { useRouter } from "next/navigation";
import axios, { AxiosError } from "axios";

interface User {
  email: string;
  role: string;
}

interface AuthResponse {
  access_token: string;
  refresh_token: string;
  access_token_expires_at: string;
  refresh_token_expires_at: string;
  user: User;
}

interface ApiResponse<T> {
  status: string;
  message: string;
  data: T;
}

interface AuthContextType {
  user: User | null;
  accessToken: string | null;
  login: (email: string, password: string) => Promise<void>;
  logout: () => Promise<void>;
  refreshAccessToken: () => Promise<AuthResponse>;
}

const AuthContext = createContext<AuthContextType>({
  user: null,
  accessToken: null,
  login: async () => { },
  logout: async () => { },
  refreshAccessToken: async () => { throw new Error("Not implemented") },
});

const API_URL = "http://localhost:8091/v1/auth";

// Create axios instance with default config
const api = axios.create({
  baseURL: API_URL,
  withCredentials: true, // Important for handling cookies
  headers: {
    'Content-Type': 'application/json',
  },
});

// Add request interceptor to add access token
api.interceptors.request.use((config) => {
  const token = (window as any).__auth_token__;
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

// Add response interceptor for error handling
api.interceptors.response.use(
  (response) => response,
  async (error: AxiosError) => {
    const originalRequest = error.config;

    // If error is 401 and we haven't tried to refresh token yet
    if (error.response?.status === 401 && originalRequest && !originalRequest.headers['retry']) {
      try {
        const response = await api.get<ApiResponse<AuthResponse>>('/refresh-token');
        const newAccessToken = response.data.data.access_token;

        // Store new access token
        (window as any).__auth_token__ = newAccessToken;

        // Retry original request with new token
        originalRequest.headers.Authorization = `Bearer ${newAccessToken}`;
        originalRequest.headers['retry'] = 'true';
        return api(originalRequest);
      } catch (refreshError) {
        // If refresh token fails, redirect to login
        window.location.href = '/login';
        return Promise.reject(refreshError);
      }
    }
    return Promise.reject(error);
  }
);

export function AuthProvider({ children }: { children: ReactNode }) {
  const [user, setUser] = useState<User | null>(null);
  const [accessToken, setAccessToken] = useState<string | null>(null);
  const router = useRouter();

  // Check authentication status on mount
  useEffect(() => {
    const checkAuth = async () => {
      try {
        const authData = await refreshAccessToken();
        setAccessToken(authData.access_token);
        (window as any).__auth_token__ = authData.access_token;
        console.log("Enter Home")
        router.push("/home");
      } catch (error) {
        console.error("Authentication check failed:", error);
        router.push("/login");
      }
    };

    checkAuth();
  }, []);

  const login = async (email: string, password: string) => {
    try {
      const response = await api.post<ApiResponse<AuthResponse>>('/login', {
        email,
        password,
      });

      const authData = response.data.data;
      setUser(authData.user);
      setAccessToken(authData.access_token);
      (window as any).__auth_token__ = authData.access_token;
      router.push("/home");
    } catch (error) {
      if (error instanceof AxiosError) {
        const message = error.response?.data?.message || "Login failed";
        throw new Error(message);
      }
      throw error;
    }
  };

  const logout = async () => {
    try {
      await api.delete('/logout');
      setUser(null);
      setAccessToken(null);
      (window as any).__auth_token__ = null;
      router.push("/login");
    } catch (error) {
      if (error instanceof AxiosError) {
        const message = error.response?.data?.message || "Logout failed";
        throw new Error(message);
      }
      throw error;
    }
  };

  const refreshAccessToken = async () => {
    try {
      const response = await api.get<ApiResponse<AuthResponse>>('/refresh-token');
      const authData = response.data.data;
      setUser(authData.user);
      setAccessToken(authData.access_token);
      return authData;
    } catch (error) {
      if (error instanceof AxiosError) {
        const message = error.response?.data?.message || "Token refresh failed";
        throw new Error(message);
      }
      throw error;
    }
  };

  return (
    <AuthContext.Provider value={{ user, accessToken, login, logout, refreshAccessToken }}>
      {children}
    </AuthContext.Provider>
  );
}

// Custom hook to use AuthContext
export function useAuth() {
  return useContext(AuthContext);
}
