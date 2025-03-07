package handler

import (
	"accountDashboardService/constants"
	reqdto "accountDashboardService/dto/request"
	respdto "accountDashboardService/dto/response"
	"encoding/json"
	"net/http"
	"strconv"
)

func (h *HandlerImpl) GetListAccount(w http.ResponseWriter, r *http.Request) error {
	functionName := "HandlerImpl.GetListAccount"

	ctx := r.Context()

	queryParams := r.URL.Query()
	customerName := queryParams.Get(constants.CustomerName)
	pageStr := queryParams.Get(constants.Page)
	countStr := queryParams.Get(constants.Count)

	pageInt, err := strconv.Atoi(pageStr)
	if err != nil {
		h.l.Debugf("[%s] = While Converting Page to Int : %s", functionName, err.Error())
		return err
	}

	countInt, err := strconv.Atoi(countStr)
	if err != nil {
		h.l.Debugf("[%s] = While Converting Count to Int : %s", functionName, err.Error())
		return err
	}

	payload := reqdto.AccountListRequest{
		CustomerName: customerName,
		Page:         pageInt,
		Count:        countInt,
	}

	result, err := h.service.GetListAccount(ctx, &payload)
	if err != nil {
		h.l.Debugf("[%s] = While Getting List Account : %s", functionName, err.Error())
		return err
	}

	res := respdto.GetListAccountResponse{
		Status:      http.StatusText(http.StatusOK),
		Message:     "Success",
		ResponseCode: strconv.Itoa(http.StatusOK),
		Data:        result.AccountList,
		Pagination: result.AccountListPagination,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(res)
}