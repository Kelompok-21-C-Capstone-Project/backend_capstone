package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controller struct {
}

func RegistrationPath(e *echo.Echo, controller Controller) {
	// API v1 basepath
	e.GET("/v1", func(c echo.Context) error {
		return c.String(http.StatusOK, "Payzone API v1.0.0 Basepath")
	})
}
