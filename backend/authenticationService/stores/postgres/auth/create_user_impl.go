package usersstore

import (
	usersreqdto "authenticationService/dto/request/auth"
	usersrespdto "authenticationService/dto/response/auth"
	cError "authenticationService/pkgs/errors"
	"context"
)

func (u *UserRepoImpl) CreateUser(ctx context.Context, payload *usersreqdto.CreateUserRequest) (*usersrespdto.UserResponse, error) {
	functionName := "UserRepoImpl.CreateUser"

	tx, err := u.DB.BeginTx(ctx, nil)
	if err != nil {
		u.l.Debugf("[%s] = While Starting Transaction : %s", functionName, err.Error())
		return nil, cError.GetError(cError.InternalServerError, err)
	}

	var userID int
	err = tx.QueryRowContext(ctx, QueryCreateUser, payload.Email, payload.Password).Scan(&userID)
	if err != nil {
		u.l.Debugf("[%s] = While Executing QueryRowContext : %s", functionName, err.Error())

		tx.Rollback()
		return nil, cError.GetError(cError.InternalServerError, err)
	}

	response := &usersrespdto.UserResponse{
		UserID: userID,
		Email:  payload.Email,
	}

	if err = tx.Commit(); err != nil {
		u.l.Debugf("[%s] = While Committing Transaction : %s", functionName, err.Error())
		return nil, cError.GetError(cError.InternalServerError, err)
	}

	return response, nil
}
