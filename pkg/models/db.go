package models

import (
	"database/sql"
	"time"
)

var db *sql.DB

func InitDB(addr string, connMaxLifetime time.Duration, maxIdleConns, maxOpenConns int) error {
	var err error
	if db, err = sql.Open("postgres", addr); err != nil {
		return err
	}
	db.SetConnMaxLifetime(connMaxLifetime)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetMaxOpenConns(maxOpenConns)
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
  			id SERIAL NOT NULL PRIMARY KEY,
  			name VARCHAR(255) NOT NULL,
  			email VARCHAR(255) NOT NULL,
  			avatar_url VARCHAR(255) NOT NULL,
  			github_id BIGINT NOT NULL,
  			admin BOOLEAN DEFAULT FALSE
		)
	`)
	return err
}
