package controllers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	jwt1 "github.com/golang-jwt/jwt/v4"
	"github.com/muhammadjon1304/jwt-authentication/cmd/models"
	"github.com/muhammadjon1304/jwt-authentication/cmd/repository"
	"net/http"
	"time"
)

type UserController struct {
	Db *sql.DB
}

const privateKey = "Secret"

func NewUserController(db *sql.DB) *UserController {
	return &UserController{
		Db: db,
	}
}

func (u *UserController) CreateUser(c *gin.Context) {
	DB := u.Db
	var user models.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	repository := repository.NewAuthRepository(DB)
	insert := repository.CreateUser(user)
	if insert {
		c.JSON(200, gin.H{"status": "success", "msg": "create user successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "failed", "msg": "insert topic failed"})
		return
	}

}

func (u *UserController) LoginUser(c *gin.Context) {
	Db := u.Db
	var db_user models.Login_User
	if err := c.ShouldBind(&db_user); err != nil {
		c.JSON(401, gin.H{"status": "failed", "msg": err})
		return
	}
	repository := repository.NewAuthRepository(Db)
	getUser := repository.LoginUser(db_user)
	if (getUser != models.User{}) {
		token := jwt1.NewWithClaims(jwt1.SigningMethodHS256, jwt1.MapClaims{
			"sub": getUser.Email,
			"exp": time.Now().Add(time.Hour * 24).Unix(),
		})
		tokenString, err := token.SignedString([]byte(privateKey))
		if err != nil {
			c.JSON(401, gin.H{"status": "failed", "msg": err})
			return
		}
		c.SetCookie("jwt", tokenString, 3600*24, "/", "localhost", false, true)

		c.Cookie("user")

		c.JSON(200, gin.H{"status": "success", "token": token, "msg": "get user successfully"})
		return
	}
}

func (u *UserController) User(c *gin.Context) {
	cookie, err := c.Cookie("jwt")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	token, err := jwt1.ParseWithClaims(cookie, &jwt1.StandardClaims{}, func(token *jwt1.Token) (interface{}, error) {
		return []byte(privateKey), nil
	})
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	claims := token.Claims.(*jwt1.StandardClaims)
	Db := u.Db
	repository := repository.NewAuthRepository(Db)
	getUser := repository.GetByEmail(claims.Issuer)
	if (getUser != models.User{}) {
		c.JSON(200, gin.H{"msg": "success", "data": getUser})
		return
	}

}
