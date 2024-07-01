package controllers

import (
	"database/sql"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/muhammadjon1304/jwt-authentication/cmd/models"
	"github.com/muhammadjon1304/jwt-authentication/cmd/repository"
	"strconv"
	"time"
)

type UserController struct {
	Db *sql.DB
}

const SecretKey = "Secret"

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
		claims := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.StandardClaims{
			Issuer:    strconv.Itoa(getUser.ID),
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		})
		token, err2 := claims.SignedString([]byte(SecretKey))
		if err2 != nil {
			c.JSON(402, gin.H{"status": "failed", "msg": "no token"})
			return
		}
		c.JSON(200, gin.H{"status": "success", "token": token, "msg": "get user successfully"})
		return
	}

	//c.SetCookie("jwt", token, 3600*24, "/", "localhost", false, true)

	//_, err4 := c.Cookie("user")
	//if err4 != nil {
	//	c.String(http.StatusNotFound, "Cookie not found")
	//	return
	//}

	//
	//c.JSON(200, gin.H{"token": token, "msg": "success"})
	//return
}
