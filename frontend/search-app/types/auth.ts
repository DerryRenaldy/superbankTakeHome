export interface User {
  email: string;
  role: string;
}

export interface AuthResponse {
  access_token: string;
  refresh_token: string;
  access_token_expires_at: string;
  refresh_token_expires_at: string;
  user: User;
}

export interface ApiResponse<T> {
  status: string;
  message: string;
  data: T;
}

export interface RefreshTokenResponse {
  access_token: string;
  access_token_expires_at: string;
  user_email: string;
}

export interface AuthContextType {
  user: string | null;
  accessToken: string | null;
  login: (email: string, password: string) => Promise<void>;
  register: (email: string, password: string) => Promise<void>;
  logout: () => Promise<void>;
  refreshAccessToken: () => Promise<ApiResponse<RefreshTokenResponse | null>>;
}
