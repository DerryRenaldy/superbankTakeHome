package usersstore

import (
	usersreqdto "authenticationService/dto/request/auth"
	usersrespdto "authenticationService/dto/response/auth"
	cError "authenticationService/pkgs/errors"
	"context"
	"errors"

	"github.com/go-sql-driver/mysql"
)

func (u *UserRepoImpl) CreateUser(ctx context.Context, payload *usersreqdto.CreateUserRequest) (*usersrespdto.UserResponse, error) {
	functionName := "UserRepoImpl.CreateUser"

	tx, err := u.DB.BeginTx(ctx, nil)
	if err != nil {
		u.l.Debugf("[%s] = While Starting Transaction : %s", functionName, err.Error())
		return nil, cError.GetError(cError.InternalServerError, err)
	}

	result, err := tx.ExecContext(ctx, QueryCreateUser, payload.Email, payload.Password)
	if err != nil {
		u.l.Debugf("[%s] = While Executing ExecContext : %s", functionName, err.Error())

		var mysqlErr *mysql.MySQLError

		if errors.As(err, &mysqlErr) {
			tx.Rollback() // Rollback the transaction on error
			return nil, err
		}

		tx.Rollback() // Rollback the transaction on error
		return nil, cError.GetError(cError.InternalServerError, err)
	}

	// Get the last inserted ID
	userID, err := result.LastInsertId()
	if err != nil {
    	tx.Rollback()
    	return nil, cError.GetError(cError.InternalServerError, err)
	}

	response := &usersrespdto.UserResponse{
		UserID: int(userID),
		Email:    payload.Email,
	}

	// Commit the transaction
	if err = tx.Commit(); err != nil {
		u.l.Debugf("[%s] = While Committing Transaction : %s", functionName, err.Error())
		return nil, cError.GetError(cError.InternalServerError, err)
	}

	return response, nil
}
