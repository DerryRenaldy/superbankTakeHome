package postgres

import (
	config "accountDashboardService/configs"
	"database/sql"
	"fmt"
	"time"

	"github.com/DerryRenaldy/logger/logger"
	_ "github.com/lib/pq" // PostgreSQL driver
)

type Connection struct {
	DBConfig config.PostgresDatabase
	log      logger.ILogger
}

func NewConnection(l logger.ILogger) *Connection {
	return &Connection{
		DBConfig: config.PostgresDatabase{
			Host:     config.Cfg.DB.Host,
			Port:     config.Cfg.DB.Port,
			Username: config.Cfg.DB.Username,
			Password: config.Cfg.DB.Password,
			DBName:   config.Cfg.DB.DBName,
		},
		log: l,
	}
}

func (db *Connection) Connect() *sql.DB {
	// PostgreSQL connection string format
	dnsAddress := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable&search_path=%s",
    db.DBConfig.Username, db.DBConfig.Password, db.DBConfig.Host, db.DBConfig.Port, db.DBConfig.DBName, "account_dashboard")

	fmt.Println(dnsAddress)

	open, err := sql.Open("postgres", dnsAddress)
	if err != nil {
		db.log.Errorf("[ERR] Error while connecting... := %v\n", err)
		return nil
	}

	// Retry connection until successful
	for open.Ping() != nil {
		db.log.Info("Attempting to connect to DB...")
		time.Sleep(5 * time.Second)
	}

	db.log.Info("Successfully connected to DB")

	return open
}
