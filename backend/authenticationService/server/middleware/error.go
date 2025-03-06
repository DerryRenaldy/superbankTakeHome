package middleware

import (
	cError "authenticationService/pkgs/errors"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type ErrHandler func(w http.ResponseWriter, r *http.Request) error

func (fn ErrHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			xerr := cError.CustomError{
				Code:    http.StatusInternalServerError,
				Status:  "Error service panic",
				Message: "Something wrong is happening, service is panicking",
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			_ = json.NewEncoder(w).Encode(xerr)
			return
		}
	}()

	if err := fn(w, r); err != nil {
		fmt.Println("Print ErrorMiddleware")

		var xerr *cError.CustomError

		if errors.As(err, &xerr) && xerr != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(xerr.Code)
			_ = json.NewEncoder(w).Encode(xerr)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			_ = json.NewEncoder(w).Encode(cError.CustomError{
				Code:    http.StatusInternalServerError,
				Status:  "Internal Server Error",
				Message: "An unexpected error occurred",
			})
		}
	}
}
