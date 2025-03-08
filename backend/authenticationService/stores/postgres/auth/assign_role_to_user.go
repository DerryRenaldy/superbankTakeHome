package usersstore

import (
	cError "authenticationService/pkgs/errors"
	"context"
	"errors"

	"github.com/lib/pq"
)

func (u *UserRepoImpl) AssignRoleToUser(ctx context.Context, userID int, roleID int) (string, error) {
	functionName := "UserRepoImpl.AssignRoleToUser"

	tx, err := u.DB.BeginTx(ctx, nil)
	if err != nil {
		u.l.Debugf("[%s] = While Starting Transaction : %s", functionName, err.Error())
		return "", cError.GetError(cError.InternalServerError, err)
	}

	_, err = tx.ExecContext(ctx, QueryAssignRoleToUser, userID, roleID)
	if err != nil {
		u.l.Debugf("[%s] = While Executing ExecContext : %s", functionName, err.Error())

		var pgErr *pq.Error

		if errors.As(err, &pgErr) {
			tx.Rollback()
			return "", err
		}

		tx.Rollback()
		return "", cError.GetError(cError.InternalServerError, err)
	}

	var roleName string

	err = tx.QueryRowContext(ctx, QueryGetRoleName, roleID).Scan(&roleName)
	if err != nil {
		u.l.Debugf("[%s] = While Executing QueryRowContext : %s", functionName, err.Error())
		tx.Rollback()
		return "", cError.GetError(cError.InternalServerError, err)
	}

	if err = tx.Commit(); err != nil {
		u.l.Debugf("[%s] = While Committing Transaction : %s", functionName, err.Error())
		return "", cError.GetError(cError.InternalServerError, err)
	}

	return roleName, nil
}
