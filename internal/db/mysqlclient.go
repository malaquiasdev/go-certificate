package db

import (
	"database/sql"
	"ekoa-certificate-generator/config"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type MySQL struct {
	connection *sql.DB
}

type IMySQL interface {
	Close() error
	GetDB() *sql.DB
}

func NewMySQLClient(cfg config.Mysql) (IMySQL, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Dbname,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %v", err)
	}

	// Test the connection
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error pinging database: %v", err)
	}

	// Set connection pool settings
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)

	return &MySQL{connection: db}, nil
}

func (m *MySQL) Close() error {
	if m.connection != nil {
		return m.connection.Close()
	}
	return nil
}

func (m *MySQL) GetDB() *sql.DB {
	return m.connection
}
