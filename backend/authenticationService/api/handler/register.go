package usershandler

import (
	usersreqdto "authenticationService/dto/request/auth"
	cError "authenticationService/pkgs/errors"
	"authenticationService/pkgs/validation"
	"encoding/json"
	"io"
	"net/http"
)

func (u *UserHandlerImpl) Register(w http.ResponseWriter, r *http.Request) error {
	functionName := "UserHandlerImpl.Register"

	ctx := r.Context()

	payload := new(usersreqdto.CreateUserRequest)

	bodyByte, err := io.ReadAll(r.Body)
	if err != nil {
		u.l.Errorf("[%s] = While Read Request Body : %s", functionName, err.Error())
		return cError.GetError(cError.InternalServerError, err)
	}

	err = json.Unmarshal(bodyByte, payload)
	if err != nil {
		u.l.Errorf("[%s] = While Unmarshal Request Body : %s", functionName, err.Error())
		return cError.GetError(cError.InternalServerError, err)
	}

	err = validation.Validate(payload)
	if err != nil {
		u.l.Errorf("[%s] = While Validate Request Body : %s", functionName, err.Error())
		return err
	}

	result, err := u.userService.Register(ctx, payload)
	if err != nil {
		return err
	}

	// Set the cookie with the token
	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    result.RefreshToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,                         // Set to true if using HTTPS
		Expires:  result.RefreshTokenExpiresAt, // Set expiration time as needed
	})

	res := struct {
		Status  int         `json:"status"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}{
		Status:  http.StatusCreated,
		Message: http.StatusText(http.StatusCreated),
		Data:    result,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	return json.NewEncoder(w).Encode(res)
}
