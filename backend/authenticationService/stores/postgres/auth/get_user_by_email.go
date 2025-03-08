package usersstore

import (
	usersrespdto "authenticationService/dto/response/auth"
	cError "authenticationService/pkgs/errors"
	"context"
	"database/sql"
	"errors"
)

func (u *UserRepoImpl) GetUserByEmail(ctx context.Context, userEmail string) (*usersrespdto.UserResponse, error) {
	functionName := "UserRepoImpl.GetUserByEmail"

	tx, err := u.DB.BeginTx(ctx, nil)
	if err != nil {
		u.l.Debugf("[%s] = While Starting Transaction : %s", functionName, err.Error())
		return nil, cError.GetError(cError.InternalServerError, err)
	}

	result := new(usersrespdto.UserResponse)

	err = tx.QueryRowContext(ctx, QueryGetOneUserByEmail, userEmail).Scan(
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
