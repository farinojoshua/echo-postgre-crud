package routes

import (
	"echo-user-app/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
)

func InitRouter() *echo.Echo {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/users", controllers.FetchUsers)
	e.GET("/users/:id", controllers.FetchUser)
	e.POST("/users/create", controllers.StoreUsers)
	e.PUT("/users/update/:id", controllers.UpdateUsers)
	e.DELETE("users/delete", controllers.DeleteUsers)

	return e
}
