package usersstore

import (
	"database/sql"

	"github.com/DerryRenaldy/logger/logger"
)

type UserRepoImpl struct {
	DB *sql.DB
	l  logger.ILogger
}

func NewUserRepoImpl(DB *sql.DB, l logger.ILogger) *UserRepoImpl {
	return &UserRepoImpl{
		DB: DB,
		l:  l,
	}
}
