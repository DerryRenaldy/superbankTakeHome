import axios, { AxiosError } from "axios";
import router from "next/router";

export const dashboardApi = axios.create({
  baseURL: process.env.NEXT_PUBLIC_DASHBOARD_API_URL,
  withCredentials: true,
  headers: {
    "Content-Type": "application/json",
  },
});

const refreshAccessToken = async () => {
  try {
    const response = await axios.get(
      `${process.env.NEXT_PUBLIC_AUTH_API_URL}/refresh-token`,
      {
        withCredentials: true,
        headers: {
          retry: "false",
        },
      }
    );

    const newAccessToken = response.data?.data?.access_token;
    if (newAccessToken) {
      localStorage.setItem("auth_token", newAccessToken);
      return newAccessToken;
    }
  } catch (error) {
    console.error("Failed to refresh token:", error);
    throw error;
  }
};

dashboardApi.interceptors.response.use(
  (response) => response,
  async (error: AxiosError) => {
    const originalRequest = error.config;
    console.log("originalRequest", originalRequest);

    if (error.response?.status === 401 && originalRequest) {
      console.warn("Unauthorized! Attempting to refresh token...");

      // Prevent retry loop
      if (!originalRequest.headers["retry"]) {
        originalRequest.headers["retry"] = "true";

        try {
          const newAccessToken = await refreshAccessToken();

          originalRequest.headers["Authorization"] = `Bearer ${newAccessToken}`;
          return dashboardApi(originalRequest);
        } catch (refreshError) {
          console.error("Refresh token failed, redirecting to login...");
        }
      }

      localStorage.removeItem("auth_token");
      router.push("/login");
    }

    return Promise.reject(error);
  }
);
