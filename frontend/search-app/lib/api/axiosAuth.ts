"use client";

import { ApiResponse, AuthResponse } from "@/types/auth";
import axios, { AxiosError } from "axios";
import router from "next/navigation";

export const authApi = axios.create({
  baseURL: process.env.NEXT_PUBLIC_AUTH_API_URL,
  withCredentials: true,
  headers: {
    "Content-Type": "application/json",
  },
});

authApi.interceptors.response.use(
  (response) => response,
  async (error: AxiosError) => {
    const originalRequest = error.config;

    // Prevent infinite loops with a retry flag
    if (
      error.response?.status === 401 &&
      originalRequest &&
      !originalRequest.headers["retry"]
    ) {
      try {
        console.log("Refreshing access token...");

        const response = await authApi.get<ApiResponse<AuthResponse>>(
          "/refresh-token"
        );
        const newAccessToken = response.data.data.access_token;

        localStorage.setItem("auth_token", newAccessToken);

        // originalRequest.headers["Authorization"] = `Bearer ${newAccessToken}`;
        originalRequest.headers["retry"] = "true"; // Prevent endless retry

        return authApi(originalRequest);
      } catch (refreshError) {
        console.error("Token refresh failed. Redirecting to login...");
        localStorage.removeItem("auth_token");
        router.redirect("/login");
        return Promise.reject(refreshError);
      }
    }

    return Promise.reject(error);
  }
);
