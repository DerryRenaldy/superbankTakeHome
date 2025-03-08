package store

import (
	reqdto "accountDashboardService/dto/request"
	respdto "accountDashboardService/dto/response"
	cError "accountDashboardService/pkgs/errors"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
)

func (r *RepoImpl) GetListAccount(ctx context.Context, req *reqdto.AccountListRequest) (*respdto.AccountListResponse, error) {
	functionName := "RepoImpl.GetListAccount"

	var whereConditions = ""
	if len(req.CustomerName) != 0 {
		whereConditions = fmt.Sprintf("AND c.name ILIKE '%%%s%%'", req.CustomerName)
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
		tx.Rollback()
		return nil, cError.GetError(cError.InternalServerError, err)
	}
	defer rows.Close()

	for rows.Next() {
		var customer respdto.Customer
		var bankAccountsJSON, pocketsJSON, termDepositsJSON sql.NullString

		if err := rows.Scan(&customer.Name, &bankAccountsJSON, &pocketsJSON, &termDepositsJSON); err != nil {
			r.l.Debugf("[%s] = While Scanning Row : %s", functionName, err.Error())
			tx.Rollback()
			return nil, cError.GetError(cError.InternalServerError, err)
		}

		if bankAccountsJSON.Valid {
			if err := json.Unmarshal([]byte(bankAccountsJSON.String), &customer.BankAccounts); err != nil {
				r.l.Debugf("[%s] = While Unmarshalling Bank Accounts : %s", functionName, err.Error())
				tx.Rollback()
				return nil, cError.GetError(cError.InternalServerError, err)
			}
		}
		if pocketsJSON.Valid {
			if err := json.Unmarshal([]byte(pocketsJSON.String), &customer.Pockets); err != nil {
				r.l.Debugf("[%s] = While Unmarshalling Pockets : %s", functionName, err.Error())
				tx.Rollback()
				return nil, cError.GetError(cError.InternalServerError, err)
			}
		}
		if termDepositsJSON.Valid {
			if err := json.Unmarshal([]byte(termDepositsJSON.String), &customer.TermDeposits); err != nil {
				r.l.Debugf("[%s] = While Unmarshalling Term Deposits : %s", functionName, err.Error())
				tx.Rollback()
				return nil, cError.GetError(cError.InternalServerError, err)
			}
		}

		result = append(result, customer)
	}

	var totalCount int
	countQuery := fmt.Sprintf(QueryGetTotalCount, whereConditions)
	if err := tx.QueryRowContext(ctx, countQuery).Scan(&totalCount); err != nil {
		r.l.Debugf("[%s] = While Querying Total Count : %s", functionName, err.Error())
		tx.Rollback()
		return nil, cError.GetError(cError.InternalServerError, err)
	}

	if err = tx.Commit(); err != nil {
		r.l.Debugf("[%s] = While Committing Transaction : %s", functionName, err.Error())
		return nil, cError.GetError(cError.InternalServerError, err)
	}

	finalResult := respdto.AccountListResponse{
		AccountListPagination: respdto.AccountListPagination{
			Page:       req.Page,
			Count:      req.Count,
			TotalCount: totalCount,
		},
		AccountList: result,
	}

	return &finalResult, nil
}
