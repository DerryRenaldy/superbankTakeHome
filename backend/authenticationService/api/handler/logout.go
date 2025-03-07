package usershandler

import (
	cError "authenticationService/pkgs/errors"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

func (u *UserHandlerImpl) Logout(w http.ResponseWriter, r *http.Request) error {
	functionName := "UserHandlerImpl.Logout"

	// Get refresh token from cookie
	refreshCookie, err := r.Cookie("refresh_token")
	if err != nil {
		u.l.Errorf("[%s] = Refresh token is required! : %s", functionName, errors.New("no refresh token provided"))
		return cError.GetError(cError.UnauthorizedError, errors.New("no refresh token provided"))
	}

	// Revoke the session in database
	if err := u.userService.Logout(r.Context(), refreshCookie.Value); err != nil {
		return err
	}

	// Clear both auth token and refresh token cookies
	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
	})


	res := struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}{
		Status:  http.StatusText(http.StatusOK),
		Message: "Successfully logged out",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(res)
}
