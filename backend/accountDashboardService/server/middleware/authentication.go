package middleware

import (
	cError "accountDashboardService/pkgs/errors"
	authagent "accountDashboardService/stores/agents/auth_agent"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/DerryRenaldy/logger/logger"
)

type AuthMiddleware struct {
	AuthClient authagent.Repository
	Logger     logger.ILogger
}

func (a *AuthMiddleware) VerifyToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authToken := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
		
		w.Header().Set("Content-Type", "application/json")

		fmt.Println("Get Header : ", authToken)

		ctx := r.Context()

		a.Logger.Infof("authToken : ", authToken)

		respVerify, err := a.AuthClient.VerifyToken(ctx, authToken)
		if err != nil {
			a.Logger.Errorf("error verifying token: %v", err)

            w.WriteHeader(http.StatusUnauthorized)
			_ = json.NewEncoder(w).Encode(cError.GetError(cError.UnauthorizedError, err))
			return
		}

        respByte,_ := json.Marshal(respVerify)

        a.Logger.Infof("respVerify : %s", respByte)

        next.ServeHTTP(w, r)
	})
}