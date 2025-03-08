package usersstore

import (
	cError "authenticationService/pkgs/errors"
	"context"
	"database/sql"
	"errors"
)

func (u *UserRepoImpl) RevokeUserSession(ctx context.Context, sessionID string) error {
	functionName := "UserRepoImpl.RevokeUserSession"	

	tx, err := u.DB.BeginTx(ctx, nil)
	if err != nil {
		u.l.Debugf("[%s] = While Starting Transaction : %s", functionName, err.Error())
		return cError.GetError(cError.InternalServerError, err)
	}

	_, err = tx.ExecContext(ctx, QueryRevokeUserSession, sessionID)
	if err != nil {
		u.l.Debugf("[%s] = While Executing ExecContext : %s", functionName, err.Error())

		if errors.Is(err, sql.ErrNoRows) {
			tx.Rollback()
			return cError.GetError(cError.InvalidRequestError, errors.New("session not found"))
		}

		tx.Rollback()
		return cError.GetError(cError.InternalServerError, err)
	}

	if err = tx.Commit(); err != nil {
		u.l.Debugf("[%s] = While Committing Transaction : %s", functionName, err.Error())
		return cError.GetError(cError.InternalServerError, err)
	}

	return nil
}
