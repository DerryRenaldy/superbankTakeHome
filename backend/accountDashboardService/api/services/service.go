package service

import (
	repository "accountDashboardService/api/repositories"
	reqdto "accountDashboardService/dto/request"
	respdto "accountDashboardService/dto/response"
	"context"

	"github.com/DerryRenaldy/logger/logger"
)

type ServiceImpl struct {
	l          logger.ILogger
	repository repository.IRepository
}

func NewServiceImpl(repository repository.IRepository, l logger.ILogger) *ServiceImpl {
	return &ServiceImpl{
		repository: repository,
		l:          l,
	}
}

type IService interface {
	GetListAccount(ctx context.Context, req *reqdto.AccountListRequest) (*respdto.AccountListResponse, error)
}

var _ IService = (*ServiceImpl)(nil)