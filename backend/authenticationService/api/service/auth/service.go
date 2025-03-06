package usersservice

import (
	usersrepo "authenticationService/api/repositories/auth"
	usersreqdto "authenticationService/dto/request/auth"
	usersrespdto "authenticationService/dto/response/auth"
	utils "authenticationService/pkgs"
	"authenticationService/pkgs/token"
	"context"

	"github.com/DerryRenaldy/logger/logger"
)

//go:generate mockgen -source service.go -destination service_mock.go -package usersservice
type UserServiceImpl struct {
	userRepo   usersrepo.IRepository
	util       utils.IUtils
	l          logger.ILogger
	tokenMaker *token.JWTImpl
}

func NewUserServiceImpl(userRepo usersrepo.IRepository, utils utils.IUtils, tokenMaker *token.JWTImpl, l logger.ILogger) *UserServiceImpl {
	return &UserServiceImpl{
		userRepo:   userRepo,
		util:       utils,
		tokenMaker: tokenMaker,
		l:          l,
	}
}

type IService interface {
	Register(ctx context.Context, payload *usersreqdto.CreateUserRequest) (*usersrespdto.RegisterLoginResponse, error)
	Login(ctx context.Context, payload *usersreqdto.LoginRequest) (*usersrespdto.RegisterLoginResponse, error)
	Logout(ctx context.Context, sessionID string) error
	RefreshToken(ctx context.Context, refreshToken string) (*usersrespdto.RenewAccessTokenResponse, error)
}

var _ IService = (*UserServiceImpl)(nil)
