package usersservice

import (
	usersrepo "authenticationService/api/repositories/auth"
	authcacherepo "authenticationService/api/repositories/auth_cache"
	config "authenticationService/configs"
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
	tokenCache authcacherepo.IRepository
	cfg        *config.Config
}

func NewUserServiceImpl(userRepo usersrepo.IRepository, utils utils.IUtils, tokenMaker *token.JWTImpl, l logger.ILogger, tokenCache authcacherepo.IRepository, cfg *config.Config) *UserServiceImpl {
	return &UserServiceImpl{
		userRepo:   userRepo,
		util:       utils,
		tokenMaker: tokenMaker,
		l:          l,
		tokenCache: tokenCache,
		cfg:        cfg,
	}
}

type IService interface {
	Register(ctx context.Context, payload *usersreqdto.CreateUserRequest) (*usersrespdto.RegisterLoginResponse, error)
	Login(ctx context.Context, payload *usersreqdto.LoginRequest) (*usersrespdto.RegisterLoginResponse, error)
	Logout(ctx context.Context, sessionID string) error
	RefreshToken(ctx context.Context, refreshToken string) (*usersrespdto.RenewAccessTokenResponse, error)
	VerifyToken(ctx context.Context, accessToken string) (*usersrespdto.VerifyTokenResponse, error)
}

var _ IService = (*UserServiceImpl)(nil)
