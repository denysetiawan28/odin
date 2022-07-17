package router

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"freyja/src/handler"
)

func InitializeRouter(server *echo.Echo, handler *handler.Handler) {
	server.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "service up")
	})

//	g := server.Group("/api")


}
