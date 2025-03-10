package usersstore

import (
	usersrespdto "authenticationService/dto/response/auth"
	cError "authenticationService/pkgs/errors"
	"context"
	"database/sql"
	"errors"
)

func (u *UserRepoImpl) GetUserById(ctx context.Context, userId int) (*usersrespdto.UserResponse, error) {
	functionName := "UserRepoImpl.GetUserById"

	tx, err := u.DB.BeginTx(ctx, nil)
	if err != nil {
		u.l.Debugf("[%s] = While Starting Transaction : %s", functionName, err.Error())
		return nil, cError.GetError(cError.InternalServerError, err)
	}

	result := new(usersrespdto.UserResponse)

	err = tx.QueryRowContext(ctx, QueryGetOneUserById, userId).Scan(
		&result.UserID, 
		&result.Role, 
		&result.Email, 
		&result.PasswordHash,
	)
	if err != nil {
		u.l.Debugf("[%s] = While Executing QueryRowContext : %s", functionName, err.Error())

		if errors.Is(err, sql.ErrNoRows) {
			tx.Rollback()
			return nil, cError.GetError(cError.InvalidRequestError, errors.New("no user found"))
		}

		tx.Rollback()
		return nil, cError.GetError(cError.InternalServerError, err)
	}

	if err = tx.Commit(); err != nil {
		u.l.Debugf("[%s] = While Committing Transaction : %s", functionName, err.Error())
		return nil, cError.GetError(cError.InternalServerError, err)
	}

	return result, nil
}
