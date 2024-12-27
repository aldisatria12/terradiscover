package handler

import (
	"net/http"

	"github.com/aldisatria12/terradiscover/apperror"
	"github.com/aldisatria12/terradiscover/dto"
	"github.com/aldisatria12/terradiscover/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ContactHandler struct {
	contactService service.ContactService
}

func NewContactHandler(s service.ContactService) ContactHandler {
	return ContactHandler{
		contactService: s,
	}
}

func (h ContactHandler) GetContact(c *gin.Context) {
	if c.Keys["user_id"] == nil {
		c.Error(apperror.NewError(nil, apperror.ErrAuthorization))
		return
	}

	reqUserId := int(c.Keys["user_id"].(float64))

	contactList, err := h.contactService.GetContact(c, reqUserId)

	if err != nil {
		c.Error(err)
		return
	}

	response := gin.H{
		"Msg":  "OK",
		"Data": contactList,
	}

	c.JSON(http.StatusOK, response)
}

func (h ContactHandler) InsertContact(c *gin.Context) {
	var newContact dto.NewContactRequest
	err := c.ShouldBindJSON(&newContact)

	if err != nil {
		c.Error(err)

		return
	}

	validate := validator.New()
	err = validate.Struct(newContact)

	if err != nil {
		c.Error(err)

		return
	}

	if c.Keys["user_id"] == nil {
		c.Error(apperror.NewError(nil, apperror.ErrAuthorization))
		return
	}

	reqUserId := int(c.Keys["user_id"].(float64))

	err = h.contactService.InsertContact(c, newContact, reqUserId)

	if err != nil {
		c.Error(err)

		return
	}

	response := gin.H{
		"Msg": "OK",
	}

	c.JSON(http.StatusOK, response)
}

func (h ContactHandler) EditContact(c *gin.Context) {
	var editContact dto.EditContactRequest
	err := c.ShouldBindJSON(&editContact)

	if err != nil {
		c.Error(err)

		return
	}

	validate := validator.New()
	err = validate.Struct(editContact)

	if err != nil {
		c.Error(err)

		return
	}

	if c.Keys["user_id"] == nil {
		c.Error(apperror.NewError(nil, apperror.ErrAuthorization))
		return
	}

	err = h.contactService.EditContact(c, editContact)

	if err != nil {
		c.Error(err)

		return
	}

	response := gin.H{
		"Msg": "OK",
	}

	c.JSON(http.StatusOK, response)
}
