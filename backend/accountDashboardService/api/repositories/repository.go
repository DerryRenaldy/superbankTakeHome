package repository

import (
	reqdto "accountDashboardService/dto/request"
	respdto "accountDashboardService/dto/response"
	store "accountDashboardService/stores/mysql"
	"context"
)

type IRepository interface {
	GetListAccount(ctx context.Context, req *reqdto.AccountListRequest) (*respdto.AccountListResponse, error)
}

var _ IRepository = (*store.RepoImpl)(nil)