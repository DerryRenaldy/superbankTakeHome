package usersrespdto

import "time"

type UserResponse struct {
	UserID       int    `json:"-"`
	Role       string    `json:"role"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"`
}

type RegisterLoginResponse struct {
	AccessToken           string        `json:"access_token"`
	RefreshToken          string        `json:"refresh_token"`
	AccessTokenExpiresAt  time.Time     `json:"access_token_expires_at"`
	RefreshTokenExpiresAt time.Time     `json:"refresh_token_expires_at"`
	User                  *UserResponse `json:"user"`
}

type Session struct {
	UserID                int    `json:"user_id"`
	RefreshToken          string    `json:"refresh_token"`
	IsRevoked             bool      `json:"is_revoked"`
	RefreshTokenExpiresAt time.Time `json:"expires_at"`
}

type RenewAccessTokenResponse struct {
	AccessToken          string    `json:"access_token"`
	AccessTokenExpiresAt time.Time `json:"access_token_expires_at"`
}

type VerifyTokenResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    VerifyTokenUserDetail `json:"data"`
}

type VerifyTokenUserDetail struct {
	Email  string `json:"email"`
	Role   string `json:"role"`
	ExpiredAt time.Time `json:"expired_at"`
}
