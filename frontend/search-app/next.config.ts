import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  publicRuntimeConfig: {
    AUTH_API_URL: "http://localhost:8091/v1/auth",
    DASHBOARD_API_URL: "http://localhost:8090/v1/dashboard",
  },
};

export default nextConfig;
