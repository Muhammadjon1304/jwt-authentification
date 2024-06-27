package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	user "github.com/muhammadjon1304/jwt-authentication/cmd/service/user"
	"log"
	"net/http"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}
func (s *APIServer) Run() error {
	router := gin.Default()
	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(router)

	log.Println("Listenig on:", s.addr)

	return http.ListenAndServe(s.addr, router)
}
