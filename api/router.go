package api

import (
	"backend_capstone/api/paymentmethod"
	"backend_capstone/api/paymentvendor"
	"backend_capstone/api/product"
	"backend_capstone/api/productbrand"
	"backend_capstone/api/productcategory"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	ProductCategory *productcategory.Controller
	ProductBrand    *productbrand.Controller
	Product         *product.Controller
	PaymentMethod   *paymentmethod.Controller
	PaymentVendor   *paymentvendor.Controller
}

func RegistrationPath(e *echo.Echo, controller Controller) {
	// API v1 basepath
	e.GET("/v1", func(c echo.Context) error {
		return c.String(http.StatusOK, "Payzone API v1.0.0 Basepath")
	})

	categoryV1 := e.Group("v1/product_categories")
	categoryV1.POST("", controller.ProductCategory.Create)
	categoryV1.GET("", controller.ProductCategory.GetAll)
	categoryV1.GET("/:id", controller.ProductCategory.GetById)
	categoryV1.PUT("/:id", controller.ProductCategory.Modify)
	categoryV1.DELETE("/:id", controller.ProductCategory.Remove)

	brandV1 := e.Group("v1/product_brands")
	brandV1.POST("", controller.ProductBrand.Create)
	brandV1.GET("", controller.ProductBrand.GetAll)
	brandV1.GET("/:id", controller.ProductBrand.GetById)
	brandV1.PUT("/:id", controller.ProductBrand.Modify)
	brandV1.DELETE("/:id", controller.ProductBrand.Remove)
	brandV1.POST("/:id/categories/:category_id", controller.ProductBrand.AddBrandCategory)
	brandV1.DELETE("/:id/categories/:category_id", controller.ProductBrand.RemoveBrandCategory)

	productV1 := e.Group("v1/products")
	productV1.POST("", controller.Product.Create)
	productV1.GET("", controller.Product.GetAll)
	productV1.GET("/clients", controller.Product.ClientGetAll)
	productV1.GET("/:id", controller.Product.GetById)
	productV1.PUT("/:id", controller.Product.Modify)
	productV1.DELETE("/:id", controller.Product.Remove)

	methodV1 := e.Group("v1/payment_methods")
	methodV1.POST("", controller.ProductCategory.Create)
	methodV1.GET("", controller.ProductCategory.GetAll)
	methodV1.GET("/:id", controller.ProductCategory.GetById)
	methodV1.PUT("/:id", controller.ProductCategory.Modify)
	methodV1.DELETE("/:id", controller.ProductCategory.Remove)

	vendorV1 := e.Group("v1/payment_vendors")
	vendorV1.POST("", controller.ProductCategory.Create)
	vendorV1.GET("", controller.ProductCategory.GetAll)
	vendorV1.GET("/:id", controller.ProductCategory.GetById)
	vendorV1.PUT("/:id", controller.ProductCategory.Modify)
	vendorV1.DELETE("/:id", controller.ProductCategory.Remove)
}
