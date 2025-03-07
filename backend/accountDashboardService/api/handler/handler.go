package handler

import (
	service "accountDashboardService/api/services"
	"net/http"

	"github.com/DerryRenaldy/logger/logger"
)

type HandlerImpl struct {
	l logger.ILogger
	service service.IService
}	

func NewHandlerImpl(log logger.ILogger, service service.IService) *HandlerImpl {
	return &HandlerImpl{
		l:       log,
		service: service,
	}
}

type IHandler interface {
	GetListAccount(w http.ResponseWriter, r *http.Request) error
}

var _ IHandler = (*HandlerImpl)(nil)
