package token

import (
	usersrespdto "authenticationService/dto/response/auth"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTImpl struct {
	secretKey string
}

func NewJWTImpl(secretKey string) *JWTImpl {
	return &JWTImpl{secretKey: secretKey}
}

func (j *JWTImpl) GenerateTokenJWT(userData *usersrespdto.UserResponse, duration time.Duration) (string, *Claims, error) {
	claims, err := NewUserClaims(userData.UserID, userData.Email, userData.Role, duration)
	if err != nil {
		return "", nil, err
	}

	fmt.Println("Token Claims Generated : ", claims)

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := newToken.SignedString([]byte(j.secretKey))
	if err != nil {
		return "", nil, err
	}

	return tokenString, claims, nil
}

func (j *JWTImpl) ValidateJWT(tokenString string) (*Claims, error) {
	newToken, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		//if token.Method.Alg() == jwt.SigningMethodHS512.Alg() {
		//	return nil, fmt.Errorf("invalid signing algorithm")
		//}

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing algorithm")
		}

		return []byte(j.secretKey), nil
	})

	if err != nil {
		// Check for specific JWT errors using errors.Is
		if errors.Is(err, jwt.ErrTokenMalformed) {
			return nil, fmt.Errorf("error: token is malformed: %w", err)
		} else if errors.Is(err, jwt.ErrTokenSignatureInvalid) {
			return nil, fmt.Errorf("error: token signature is invalid: %w", err)
		} else if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, fmt.Errorf("error: token is expired: %w", err)
		} else if errors.Is(err, jwt.ErrTokenNotValidYet) {
			return nil, fmt.Errorf("error: token is not valid yet: %w", err)
		} else if errors.Is(err, jwt.ErrTokenUnverifiable) {
			return nil, fmt.Errorf("error: token is unverifiable: %w", err)
		} else {
			// Handle any other errors
			return nil, fmt.Errorf("some error occurred: %w", err)
		}
	}

	if !newToken.Valid {
		return nil, fmt.Errorf("error token is invalid")
	}

	claims, ok := newToken.Claims.(*Claims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}
	return claims, nil
}
