package mysql

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/maiquocthinh/go-comic/config"
)

const (
	maxOpenConns    = 6
	maxIdleConns    = 4
	connMaxLifetime = 90 * time.Second // 1.5 minutes
	connMaxIdleTime = 60 * time.Second // 1 minutes
)

func NewMysqlDB(cfg *config.MySQLConfig) (*sqlx.DB, error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
	)

	db, err := sqlx.Open(cfg.DriverName, connectionString)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxLifetime(connMaxLifetime)
	db.SetConnMaxIdleTime(connMaxIdleTime)

	return db, nil
}
