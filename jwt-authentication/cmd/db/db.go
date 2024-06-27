package db

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"log"
)

func NewPostgreSQLStorage(cfg mysql.Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	return db, nil
}
