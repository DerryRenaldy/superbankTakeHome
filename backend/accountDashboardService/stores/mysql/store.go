package store

import (
	"database/sql"

	"github.com/DerryRenaldy/logger/logger"
)

type RepoImpl struct {
	DB *sql.DB
	l  logger.ILogger
}

func NewRepoImpl(DB *sql.DB, l logger.ILogger) *RepoImpl {
	return &RepoImpl{
		DB: DB,
		l:  l,
	}
}
