"use client"

import { createContext, useContext, useState, ReactNode, useEffect } from "react";
import { useRouter } from "next/navigation";
import { AxiosError } from "axios";
import { AuthContextType, User, AuthResponse, ApiResponse, RefreshTokenResponse } from "@/types/auth";
import { authApi } from "@/lib/api/axiosAuth";

const AuthContext = createContext<AuthContextType | null>(null);

export function AuthProvider({ children }: { children: ReactNode }) {
  const [user, setUser] = useState<string | null>(null);
  const [accessToken, setAccessToken] = useState<string | null>(null);
  const router = useRouter();

  useEffect(() => {
    const checkAuth = async () => {
      const storedToken = localStorage.getItem("auth_token");

      if (!storedToken) {
        console.warn("No auth token found. Redirecting to login.");
        return router.push("/login");
      }

      try {
        const authData = await refreshAccessToken();
        console.log("authData", authData);
        setAccessToken(authData.data.access_token);
      } catch (error) {
        console.error("Authentication check failed:", error);
        router.push("/login");
      }
    };

    checkAuth();
  }, []);

  const register = async (email: string, password: string) => {
    try {
      const response = await authApi.post<ApiResponse<AuthResponse>>('/register', {
        email,
        password,
      });

      const authData = response.data.data;
      setUser(authData.user.email);
      setAccessToken(authData.access_token);
      localStorage.setItem('auth_token', authData.access_token);
      router.push("/home");
    } catch (error) {
      if (error instanceof AxiosError) {
        const message = error.response?.data?.message || "Register failed";
        throw new Error(message);
      }
      throw error;
    }
  };

  const login = async (email: string, password: string) => {
    try {
      const response = await authApi.post<ApiResponse<AuthResponse>>('/login', {
        email,
        password,
      });

      const authData = response.data.data;
      setUser(authData.user.email);
      setAccessToken(authData.access_token);
      localStorage.setItem('auth_token', authData.access_token);
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
      await authApi.delete('/logout');
      setUser(null);
      setAccessToken(null);
      localStorage.removeItem('auth_token');
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
      const response = await authApi.get<ApiResponse<RefreshTokenResponse>>('/refresh-token', {
        withCredentials: true,
        headers: {
          "retry": "false"
        }
      });

      const authResponseData = response.data;
      setUser(authResponseData.data.user_email);
      console.log("user original data", user);
      console.log("user", user);
      setAccessToken(authResponseData.data.access_token);
      return authResponseData;
    } catch (error) {
      if (error instanceof AxiosError) {
        const message = error.response?.data?.message || "Token refresh failed";
        throw new Error(message);
      }
      throw error;
    }
  };

  return (
    <AuthContext.Provider value={{ user, accessToken, register, login, logout, refreshAccessToken }}>
      {children}
    </AuthContext.Provider>
  );
}

export function useAuth() {
  const context = useContext(AuthContext);
  if (!context) {
    throw new Error("useAuth must be used within a AuthProvider");
  }

  return context;
}
