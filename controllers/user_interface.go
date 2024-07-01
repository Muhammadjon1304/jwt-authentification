package controllers

import "github.com/gin-gonic/gin"

type UserInterface interface {
	CreateUser(ctx *gin.Context)
	Login(ctx *gin.Context)
}
