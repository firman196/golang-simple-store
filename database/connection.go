package database

import (
	"database/sql"
	"time"
)

func Database(db_host string, db_port string, db_name string) (*sql.DB, error) {
	db, err := sql.Open("mysql", "root@tcp("+db_host+":"+db_port+")/"+db_name)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(25)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db, err
}
