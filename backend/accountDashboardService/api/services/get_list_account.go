package service

import (
	reqdto "accountDashboardService/dto/request"
	respdto "accountDashboardService/dto/response"
	"context"
)

func (s *ServiceImpl) GetListAccount(ctx context.Context, req *reqdto.AccountListRequest) (*respdto.AccountListResponse, error) {
	functionName := "ServiceImpl.GetListAccount"

	result, err := s.repository.GetListAccount(ctx, req)
	if err != nil {
		s.l.Debugf("[%s] = While Getting List Account : %s", functionName, err.Error())
		return nil, err
	}

	return result, nil
}