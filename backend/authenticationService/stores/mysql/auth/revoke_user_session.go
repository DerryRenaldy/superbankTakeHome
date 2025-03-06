package usersstore

import (
	cError "authenticationService/pkgs/errors"
	"context"
	"errors"

	"github.com/go-sql-driver/mysql"
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

		var mysqlErr *mysql.MySQLError

		if errors.As(err, &mysqlErr) {
			tx.Rollback() // Rollback the transaction on error
			return err
		}

		tx.Rollback() // Rollback the transaction on error
		return cError.GetError(cError.InternalServerError, err)
	}

	// Commit the transaction
	if err = tx.Commit(); err != nil {
		u.l.Debugf("[%s] = While Committing Transaction : %s", functionName, err.Error())
		return cError.GetError(cError.InternalServerError, err)
	}

	return nil
}
