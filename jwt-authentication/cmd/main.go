package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	controllers "github.com/muhammadjon1304/jwt-authentication/cmd/controllers"
	"github.com/muhammadjon1304/jwt-authentication/cmd/initializers"
)

func main() {
	db := initializers.ConnectDB()
	router := gin.Default()
	controller := controllers.NewUserController(db)
	router.POST("/register", controller.CreateUser)
	router.POST("/login", controller.LoginUser)
	router.Run(":8080")

	router.Use(cors.New(cors.Config{
		AllowFiles: true,
	}))
}
