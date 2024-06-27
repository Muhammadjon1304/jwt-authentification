package main

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/muhammadjon1304/jwt-authentication/cmd/api"
	"github.com/muhammadjon1304/jwt-authentication/cmd/config"
	"github.com/muhammadjon1304/jwt-authentication/cmd/db"
	"log"
)

func main() {
	db, err := db.NewPostgreSQLStorage(mysql.Config{
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

	initStorage(db)

	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected succefully")
}
