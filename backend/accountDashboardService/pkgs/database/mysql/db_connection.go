package mysql

import (
	config "accountDashboardService/configs"
	"database/sql"
	"fmt"
	"time"

	"github.com/DerryRenaldy/logger/logger"
	_ "github.com/go-sql-driver/mysql"
)

type Connection struct {
	DBConfig config.MySQLDatabase
	log      logger.ILogger
}

func NewConnection(l logger.ILogger) *Connection {
	return &Connection{
		DBConfig: config.MySQLDatabase{
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
	dnsAddress := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", db.DBConfig.Username, db.DBConfig.Password, db.DBConfig.Host, db.DBConfig.Port, db.DBConfig.DBName)

	open, err := sql.Open("mysql", dnsAddress)
	if err != nil {
		db.log.Errorf("[ERR] Error while connecting... := %v\n", err)
		return nil
	}

	for open.Ping() != nil {
		db.log.Info("Attempting connect to DB...")
		time.Sleep(5 * time.Second)
	}

	db.log.Info("Successfully connected to DB")

	return open
}
