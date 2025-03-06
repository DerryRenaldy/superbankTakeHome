package usersreqdto

import "time"

// {
// "username" : "derryrenaldy2",
// "email" : "derryrenaldy2@gmail.com",
// "password": "test123"
// }
type CreateUserRequest struct {
	UserID   int    `json:"user_id"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// {
// "user_identity" : "derryrenaldy2",
// "password": "test123"
// }
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password     string `json:"password" validate:"required"`
}

type CreateSessionRequest struct {
	SessionID             string    `json:"session_id"`
	UserID                string    `json:"user_id"`
	RefreshToken          string    `json:"refresh_token"`
	RefreshTokenExpiresAt time.Time `json:"refresh_token_expires_at"`
}

type RenewAccessTokenRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}
