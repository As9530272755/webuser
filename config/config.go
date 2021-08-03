package config

import (
	"database/sql"
)

var DB *sql.DB

func ConnectDB(driverName, dsn string) error {
	var err error
	DB, err = sql.Open(driverName, dsn)
	if err != nil {
		return err
	}

	if err := DB.Ping(); err != nil {
		return err
	}
	return nil
}

func CloseDB() error {
	err := DB.Close()
	if err != nil {
		return err
	}
	return nil
}
