package handler

import (
	"net/http"

	"github.com/aldisatria12/terradiscover/apperror"
	"github.com/aldisatria12/terradiscover/service"
	"github.com/gin-gonic/gin"
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
