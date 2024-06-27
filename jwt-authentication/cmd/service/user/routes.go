package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/muhammadjon1304/jwt-authentication/cmd/auth"
	"github.com/muhammadjon1304/jwt-authentication/cmd/types"
	"github.com/muhammadjon1304/jwt-authentication/cmd/utils"
	"log"
	"net/http"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) RegisterRoutes(router *gin.Engine) {
	router.POST("/login", h.handleLogin)
	router.POST("/register", h.handleRegister)
}
func (h *Handler) handleLogin(c *gin.Context) {

}

func (h *Handler) handleRegister(c *gin.Context) {
	var payload types.RegisterUserPayload
	if err := utils.ParseJSON(c.Request, payload); err != nil {
		utils.WriteError(c.Writer, http.StatusBadRequest, err)
		return
	}
	if err := utils.Validate.Struct(payload); err != nil {
		error := err.(validator.ValidationErrors)
		utils.WriteError(c.Writer, http.StatusBadRequest, fmt.Errorf("invalid payload %v", error))
		return
	}

	_, err := h.store.GetUserByEmail(payload.Email)
	if err == nil {
		utils.WriteError(c.Writer, http.StatusBadRequest, fmt.Errorf("User with email %s already exists"))
		return
	}

	hashedPassword, err := auth.HasPassword(payload.Password)

	if err != nil {
		log.Fatal(err)
	}

	err = h.store.CreateUser(types.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  hashedPassword,
	})
	if err == nil {
		utils.WriteError(c.Writer, http.StatusInternalServerError, err)
		return
	}

	utils.WriteError(c.Writer, http.StatusCreated, err)
}
