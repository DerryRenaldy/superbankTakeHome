package usershandler

import (
	usersservice "authenticationService/api/service/auth"
	"net/http"

	"github.com/DerryRenaldy/logger/logger"
)

type UserHandlerImpl struct {
	userService usersservice.IService
	l           logger.ILogger
}

func NewUserHandlerImpl(userService usersservice.IService, l logger.ILogger) *UserHandlerImpl {
	return &UserHandlerImpl{
		userService: userService,
		l:           l,
	}
}

type IHandler interface {
	Register(w http.ResponseWriter, r *http.Request) error
	Login(w http.ResponseWriter, r *http.Request) error
	Logout(w http.ResponseWriter, r *http.Request) error
	RefreshToken(w http.ResponseWriter, r *http.Request) error
}

var _ IHandler = (*UserHandlerImpl)(nil)
