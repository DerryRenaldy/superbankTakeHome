package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Claims struct {
	UserID   int `json:"user_id"`
	Email    string `json:"email"`
	Role 	 string `json:"role"`
	jwt.RegisteredClaims
}

func NewUserClaims(userID int, email string, role string, duration time.Duration) (*Claims, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, fmt.Errorf("error generating token ID: %v", err)
	}

	return &Claims{
		UserID:   userID,
		Email:    email,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        tokenID.String(),
			Subject:   email,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
		},
	}, nil
}

func (u *Claims) Valid() error {
	// Manually check the expiration time
	fmt.Println("Token expired at :", u.ExpiresAt)
	if u.ExpiresAt == nil || u.ExpiresAt != nil && u.ExpiresAt.Before(time.Now()) {
		return fmt.Errorf("error: token has expired")
	}

	return nil
}
