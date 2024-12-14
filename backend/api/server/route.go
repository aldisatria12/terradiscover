package server

import (
	"database/sql"
	"net/http"

	"github.com/aldisatria12/terradiscover/handler"
	"github.com/aldisatria12/terradiscover/repository"
	"github.com/aldisatria12/terradiscover/service"
	"github.com/gin-gonic/gin"
)

type Route struct {
	userHandler *handler.UserHandler
}

func NewRoute(db *sql.DB) Route {
	dataStore := repository.NewDataStore(db)
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(dataStore, userRepository)
	userHandler := handler.NewUserHandler(userService)

	return Route{
		userHandler: &userHandler,
	}
}

func (route *Route) SetRoutes() http.Handler {
	r := gin.New()
	r.POST("/user/login", route.userHandler.Login)
	r.POST("/user/register", route.userHandler.Register)

	return r
}
