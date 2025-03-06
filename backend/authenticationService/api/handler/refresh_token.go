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

	refreshToken := r.URL.Query().Get("token")
	if refreshToken == "" {
		u.l.Errorf("[%s] = Refresh token is required! : %s", functionName, errors.New("refresh token is required"))
		return cError.GetError(cError.BadRequestError, errors.New("refresh token is required"))
	}

	result, err := u.userService.RefreshToken(ctx, refreshToken)
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
