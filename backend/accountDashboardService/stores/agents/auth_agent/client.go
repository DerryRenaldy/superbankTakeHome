package authagent

import (
	config "accountDashboardService/configs"
	respdto "accountDashboardService/dto/response"
	"context"

	"github.com/DerryRenaldy/logger/logger"
)

type (
	Repository interface {
		VerifyToken(ctx context.Context, accessToken string) (*respdto.VerifyTokenResponse, error)
	}

	authClientImpl struct {
		l     logger.ILogger
		cfg     *config.Config
	}
)

func NewClientAuth(cfg *config.Config, log logger.ILogger) *authClientImpl {
	return &authClientImpl{
		cfg:     cfg,
		l:     log,
	}
}
