package middleware

import (
	"net/http"
	"time"

	"github.com/aldisatria12/terradiscover/apperror"
	"github.com/aldisatria12/terradiscover/util/logger"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log := logger.Log
		startTime := time.Now()
		ctx.Next()
		endTime := time.Now()

		latencyTime := endTime.Sub(startTime)
		reqMethod := ctx.Request.Method
		reqUri := ctx.Request.RequestURI
		statusCode := ctx.Writer.Status()
		if len(ctx.Errors) > 0 {
			err := ctx.Errors[0]
			switch e := err.Err.(type) {
			case apperror.CustomError:
				statusCode = e.StatusCode
			case validator.ValidationErrors:
				statusCode = http.StatusBadRequest
			default:
				statusCode = http.StatusInternalServerError
			}
		}
		clientIP := ctx.ClientIP()

		if lastErr := ctx.Errors.Last(); lastErr != nil {
			log.WithFields(map[string]any{
				"METHOD":    reqMethod,
				"URI":       reqUri,
				"STATUS":    statusCode,
				"LATENCY":   latencyTime,
				"CLIENT_IP": clientIP,
			}).Error()
			return
		}

		log.WithFields(map[string]any{
			"METHOD":    reqMethod,
			"URI":       reqUri,
			"STATUS":    statusCode,
			"LATENCY":   latencyTime,
			"CLIENT_IP": clientIP,
		}).Infof("REQUEST %s %s SUCCESS", reqMethod, reqUri)
	}
}
