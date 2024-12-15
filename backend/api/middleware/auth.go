package middleware

import (
	"strings"

	"github.com/aldisatria12/terradiscover/apperror"
	"github.com/aldisatria12/terradiscover/util"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if token == "" {
			c.Error(apperror.NewError(nil, apperror.ErrAuthorization))
			return
		}

		splitted := strings.Split(token, " ")
		if len(splitted) != 2 {
			c.Error(apperror.NewError(nil, apperror.ErrAuthorization))
			return
		}

		claims, err := util.ParseAndVerify(splitted[1])
		if err != nil {
			c.Error(apperror.NewError(nil, apperror.ErrAuthorization))
			return
		}
		c.Set("user_id", claims["user_id"])
		c.Next()
	}
}
