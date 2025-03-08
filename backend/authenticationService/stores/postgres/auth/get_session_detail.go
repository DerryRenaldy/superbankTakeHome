package usersstore

import (
	usersrespdto "authenticationService/dto/response/auth"
	cError "authenticationService/pkgs/errors"
	"context"
	"database/sql"
	"errors"
)

func (u *UserRepoImpl) GetSessionDetail(ctx context.Context, refreshToken string) (*usersrespdto.Session, error) {
	functionName := "UserRepoImpl.GetSessionDetail"

	tx, err := u.DB.BeginTx(ctx, nil)
	if err != nil {
		u.l.Debugf("[%s] = While Starting Transaction : %s", functionName, err.Error())
		return nil, cError.GetError(cError.InternalServerError, err)
	}

	result := new(usersrespdto.Session)

	err = tx.QueryRowContext(ctx, QueryGetSessionDetail, refreshToken).Scan(
		&result.UserID, 
		&result.RefreshToken, 
		&result.IsRevoked, 
		&result.RefreshTokenExpiresAt,
	)
	if err != nil {
		u.l.Debugf("[%s] = While Executing QueryRowContext : %s", functionName, err.Error())

		if errors.Is(err, sql.ErrNoRows) {
			tx.Rollback()
			return nil, cError.GetError(cError.InvalidRequestError, errors.New("no session found"))
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
