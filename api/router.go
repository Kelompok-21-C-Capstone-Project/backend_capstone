package api

import (
	"backend_capstone/api/payments"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	PaymentV1Controller *payments.Controller
}

func RegistrationPath(e *echo.Echo, controller Controller) {
	// API v1 basepath
	e.GET("/v1", func(c echo.Context) error {
		return c.String(http.StatusOK, "Payzone API v1.0.0 Basepath")
	})

	paymentV1 := e.Group("/v1/payments")
	paymentV1.GET("", controller.PaymentV1Controller.Create)
	paymentV1.GET("/finish", controller.PaymentV1Controller.Modify)
}
