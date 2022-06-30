package api

import (
	"backend_capstone/api/user"
	"backend_capstone/middleware"
	"net/http"

	"github.com/labstack/echo"
)

type Controller struct {
}

func Init(e *echo.Echo) {
	e.GET("/generate-hash/:password", user.GenerateHashPassword)
}

func RegistrationPath(e *echo.Echo, controller Controller) {
	// API v1 basepath
	e.GET("/v1", func(c echo.Context) error {
		return c.String(http.StatusOK, "Payzone API v1.0.0 Basepath")
	})

	e.GET("/user", user.GetAllData, middleware.IsAuthenticated)
}
