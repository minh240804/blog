package model

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123"
	dbname   = "DemoBlog"
)

var DB *sql.DB

func DbConnect() error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
		defer db.Close()
		return err
	}
	DB = db
	return nil

}
