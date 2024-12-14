package handler

import (
	"net/http"

	"github.com/aldisatria12/terradiscover/dto"
	"github.com/aldisatria12/terradiscover/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(s service.UserService) UserHandler {
	return UserHandler{
		userService: s,
	}
}

func (h UserHandler) Login(c *gin.Context) {
	var logUser dto.UserLoginRequest
	err := c.ShouldBindJSON(&logUser)

	if err != nil {
		c.Error(err)

		return
	}

	validate := validator.New()
	err = validate.Struct(logUser)

	if err != nil {
		c.Error(err)

		return
	}

	user, err := h.userService.Login(c, logUser)

	if err != nil {
		c.Error(err)

		return
	}

	response := gin.H{
		"Msg":  "OK",
		"Data": user,
	}

	c.JSON(http.StatusOK, response)
}

func (h UserHandler) Register(c *gin.Context) {
	var newUser dto.UserRegisterRequest
	err := c.ShouldBindJSON(&newUser)

	if err != nil {
		c.Error(err)

		return
	}

	validate := validator.New()
	err = validate.Struct(newUser)

	if err != nil {
		c.Error(err)

		return
	}

	result, err := h.userService.Register(c, newUser)

	if err != nil {
		c.Error(err)

		return
	}

	response := gin.H{
		"Msg":  "OK",
		"Data": result,
	}

	c.JSON(http.StatusOK, response)
}
