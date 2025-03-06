package middleware

import (
	utils "authenticationService/pkgs"
	cError "authenticationService/pkgs/errors"
	"authenticationService/pkgs/token"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// writeJSONError writes a JSON error response
func writeJSONError(w http.ResponseWriter, err *cError.CustomError) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(err.Code)
    json.NewEncoder(w).Encode(err)
}

// AuthMiddleware returns a middleware that verifies JWT tokens
func AuthMiddleware() mux.MiddlewareFunc {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            cookie, err := r.Cookie("AuthToken")
            if err != nil {
                log.Printf("Error Getting Cookie : %s", err.Error())
                writeJSONError(w, &cError.CustomError{
                    Code:    http.StatusUnauthorized,
                    Status:  "Unauthorized",
                    Message: "Missing or invalid AuthToken",
                })
                return
            }

            tokenString := cookie.Value
            claims := &token.Claims{}

            claims, err = utils.ValidateJWT(tokenString)
            if err != nil {
                fmt.Printf("Error ValidateJWT : %s", err.Error())
                writeJSONError(w, &cError.CustomError{
                    Code:    http.StatusUnauthorized,
                    Status:  "Unauthorized",
                    Message: "Invalid token",
                })
                return
            }

            err = claims.Valid()
            if err != nil {
                writeJSONError(w, &cError.CustomError{
                    Code:    http.StatusUnauthorized,
                    Status:  "Unauthorized",
                    Message: "Token Expired",
                })
                return
            }

			type contextKey string
			const UserClaimsKey contextKey = "userClaims"

            // Set claims in request context
			ctx := r.Context()
			ctx = context.WithValue(ctx, UserClaimsKey, claims)
            next.ServeHTTP(w, r.WithContext(ctx))
        })
    }
}