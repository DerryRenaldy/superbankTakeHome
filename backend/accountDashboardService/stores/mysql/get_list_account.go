package store

import (
	reqdto "accountDashboardService/dto/request"
	respdto "accountDashboardService/dto/response"
	cError "accountDashboardService/pkgs/errors"
	"context"
	"encoding/json"
	"fmt"
)

func (r *RepoImpl) GetListAccount(ctx context.Context, req *reqdto.AccountListRequest) (*respdto.AccountListResponse, error) {
	functionName := "RepoImpl.GetListAccount"

	whereConditions := ""

	if len(req.CustomerName) != 0 {
		whereConditions += fmt.Sprintf(" WHERE c.name LIKE ('%s')", req.CustomerName)
	}

	// Validate pagination parameters
	if req.Page <= 0 || req.Count <= 0 {
		return nil, cError.GetError(cError.BadRequestError, fmt.Errorf("invalid pagination parameters"))
	}

	offset := (req.Page - 1) * req.Count

	queryString := fmt.Sprintf(QueryGetListAccount, whereConditions, req.Count, offset)

	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		r.l.Debugf("[%s] = While Starting Transaction : %s", functionName, err.Error())
		return nil, cError.GetError(cError.InternalServerError, err)
	}

	var result []respdto.Customer

	rows, err := tx.QueryContext(ctx, queryString)
	if err != nil {
		r.l.Debugf("[%s] = While Querying Accounts : %s", functionName, err.Error())
		return nil, cError.GetError(cError.InternalServerError, err)
	}
	defer rows.Close()

	for rows.Next() {
		var customer respdto.Customer
		var bankAccountsJSON, pocketsJSON, termDepositsJSON string

		if err := rows.Scan(&customer.Name, &bankAccountsJSON, &pocketsJSON, &termDepositsJSON); err != nil {
			r.l.Debugf("[%s] = While Scanning Row : %s", functionName, err.Error())
			return nil, cError.GetError(cError.InternalServerError, err)
		}

		// Unmarshal JSON data into the struct
		if err := json.Unmarshal([]byte(bankAccountsJSON), &customer.BankAccounts); err != nil {
			r.l.Debugf("[%s] = While Unmarshalling Bank Accounts : %s", functionName, err.Error())
			return nil, cError.GetError(cError.InternalServerError, err)
		}
		if err := json.Unmarshal([]byte(pocketsJSON), &customer.Pockets); err != nil {
			r.l.Debugf("[%s] = While Unmarshalling Pockets : %s", functionName, err.Error())
			return nil, cError.GetError(cError.InternalServerError, err)
		}
		if err := json.Unmarshal([]byte(termDepositsJSON), &customer.TermDeposits); err != nil {
			r.l.Debugf("[%s] = While Unmarshalling Term Deposits : %s", functionName, err.Error())
			return nil, cError.GetError(cError.InternalServerError, err)
		}

		result = append(result, customer)
	}

	// Query to get the total count of accounts
	var totalCount int
	countQuery := fmt.Sprintf(QueryGetTotalCount, whereConditions)
	if err := tx.QueryRowContext(ctx, countQuery).Scan(&totalCount); err != nil {
		r.l.Debugf("[%s] = While Querying Total Count : %s", functionName, err.Error())
		return nil, cError.GetError(cError.InternalServerError, err)
	}

	// Commit the transaction
	if err = tx.Commit(); err != nil {
		r.l.Debugf("[%s] = While Committing Transaction : %s", functionName, err.Error())
		return nil, cError.GetError(cError.InternalServerError, err)
	}

	finalResult := respdto.AccountListResponse{}
	finalResult.AccountListPagination.Page = req.Page
	finalResult.AccountListPagination.Count = req.Count
	finalResult.AccountListPagination.TotalCount = totalCount
	finalResult.AccountList = result

	return &finalResult, nil
}
