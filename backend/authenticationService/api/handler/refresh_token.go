package usershandler

import (
	cError "authenticationService/pkgs/errors"
	"encoding/json"
	"errors"
	"net/http"
)

func (u *UserHandlerImpl) RefreshToken(w http.ResponseWriter, r *http.Request) error {
	functionName := "UserHandlerImpl.RefreshToken"

	ctx := r.Context()

	// Get refresh token from cookie
	refreshToken, err := r.Cookie("refresh_token")
	if err != nil {
		u.l.Errorf("[%s] = Refresh token is required! : %s", functionName, errors.New("refresh token is required"))
		return cError.GetError(cError.UnauthorizedError, errors.New("no refresh token provided"))
	}

	result, err := u.userService.RefreshToken(ctx, refreshToken.Value)
	if err != nil {
		return err
	}

	res := struct {
		Status  string      `json:"status"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}{
		Status:  http.StatusText(http.StatusOK),
		Message: "Success",
		Data:    result,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(res)
}
