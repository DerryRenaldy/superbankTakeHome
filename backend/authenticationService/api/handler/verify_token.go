package usershandler

import (
	cError "authenticationService/pkgs/errors"
	"encoding/json"
	"errors"
	"net/http"
)

func (u *UserHandlerImpl) VerifyToken(w http.ResponseWriter, r *http.Request) error {
	functionName := "UserHandlerImpl.VerifyToken"

	ctx := r.Context()
	
	accessToken := r.URL.Query().Get("access_token")
	if accessToken == "" {
		u.l.Errorf("[%s] = Access token is required! : %s", functionName, errors.New("access token is required"))
		return cError.GetError(cError.UnauthorizedError, errors.New("no access token provided"))
	}

	// Verify refresh token
	result, err := u.userService.VerifyToken(ctx, accessToken)
	if err != nil {
		return err
	}

	// Return response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(result)
}


