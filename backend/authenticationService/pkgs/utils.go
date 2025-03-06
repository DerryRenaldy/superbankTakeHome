package utils

import (
	"authenticationService/pkgs/token"
	"errors"
	"fmt"
	"regexp"

	"github.com/DerryRenaldy/logger/logger"

	config "authenticationService/configs"
	usersrespdto "authenticationService/dto/response/auth"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UtilsImpl struct {
	l logger.ILogger
}

//go:generate mockgen -source=utils.go -destination=utils_mock.go -package=utils
type IUtils interface {
	GeneratePasswordHash(password string) (string, error)
	MatchPassword(password string, hash string) bool
	GenerateTokenJWT(userData *usersrespdto.UserResponse, duration time.Duration) (string, *token.Claims, error)
	ValidateJWT(tokenString string) (*token.Claims, error)
}

var _ IUtils = (*UtilsImpl)(nil)

func NewUtilsImpl(l logger.ILogger) *UtilsImpl {
	return &UtilsImpl{l: l}
}

func (u *UtilsImpl) MatchPassword(password string, hash string) bool {
	// Original implementation
	return MatchPassword(password, hash)
}

func (u *UtilsImpl) GenerateTokenJWT(userData *usersrespdto.UserResponse, duration time.Duration) (string, *token.Claims, error) {
	// Original implementation
	return GenerateTokenJWT(userData, duration)
}

func (u *UtilsImpl) GeneratePasswordHash(password string) (string, error) {
	return GeneratePasswordHash(password)
}

func (u *UtilsImpl) ValidateJWT(tokenString string) (*token.Claims, error) {
	return ValidateJWT(tokenString)
}

func GeneratePasswordHash(password string) (string, error) {
	password = strings.Trim(password, " ")
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func MatchPassword(password string, hash string) bool {
	password = strings.Trim(password, " ")
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}

func GenerateTokenJWT(userData *usersrespdto.UserResponse, duration time.Duration) (string, *token.Claims, error) {
	claims, err := token.NewUserClaims(userData.UserID, userData.Email, userData.Role, duration)
	if err != nil {
		return "", nil, err
	}

	fmt.Println("Token Claims Generated : ", claims)

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := newToken.SignedString([]byte(config.Cfg.JWTSecret))
	if err != nil {
		return "", nil, err
	}

	return tokenString, claims, nil
}

func ValidateJWT(tokenString string) (*token.Claims, error) {
	newToken, err := jwt.ParseWithClaims(tokenString, &token.Claims{}, func(token *jwt.Token) (interface{}, error) {
		//if token.Method.Alg() == jwt.SigningMethodHS512.Alg() {
		//	return nil, fmt.Errorf("invalid signing algorithm")
		//}

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing algorithm")
		}

		return []byte(config.Cfg.JWTSecret), nil
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

	claims, ok := newToken.Claims.(*token.Claims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}
	return claims, nil
}

func IsInputEmail(data string) bool {
	regex := "[\\w\\.-_]+@[\\w\\.-]+\\.\\w+"

	// Compile the regular expression
	re := regexp.MustCompile(regex)

	// Use the MatchString method to check if the data matches the email pattern
	return re.MatchString(data)
}
