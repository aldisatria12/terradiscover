package middleware

import (
	"net/http"

	"github.com/aldisatria12/terradiscover/apperror"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ErrorMiddleware(c *gin.Context) {
	c.Next()
	if len(c.Errors) > 0 {
		err := c.Errors[0]
		switch e := err.Err.(type) {
		case apperror.CustomError:
			c.AbortWithStatusJSON(e.StatusCode, gin.H{"message": e})
		case validator.ValidationErrors:
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": apperror.ErrBinding})
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": apperror.ErrServer})
		}
	}
}
