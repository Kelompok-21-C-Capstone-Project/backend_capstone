package api

import (
	"backend_capstone/api/productcategory"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	ProductCategory *productcategory.Controller
}

func RegistrationPath(e *echo.Echo, controller Controller) {
	// API v1 basepath
	e.GET("/v1", func(c echo.Context) error {
		return c.String(http.StatusOK, "Payzone API v1.0.0 Basepath")
	})

	categoryV1 := e.Group("v1/categories")
	categoryV1.POST("", controller.ProductCategory.Create)
	categoryV1.GET("", controller.ProductCategory.GetAll)
	categoryV1.GET("/:id", controller.ProductCategory.GetById)
	categoryV1.PUT("/:id", controller.ProductCategory.Modify)
	categoryV1.DELETE("/:id", controller.ProductCategory.Remove)
}
