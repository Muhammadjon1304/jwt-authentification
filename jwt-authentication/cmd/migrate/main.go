package main

import (
	"github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/muhammadjon1304/jwt-authentication/cmd/config"
	"github.com/muhammadjon1304/jwt-authentication/cmd/db"
	"log"
	"os"
)

func main() {
	_, err := db.NewPostgreSQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.New(
		"file://cmd/migrate/migrations",
		"postgres",
	)
	if err != nil {
		log.Fatal(err)
	}
	cmd := os.Args[len(os.Args)-1]

	if cmd == "up" {
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}

	if cmd == "down" {
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}
}
