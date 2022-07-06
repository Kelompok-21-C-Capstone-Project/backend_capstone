package api

import (
	AdminMiddleware "backend_capstone/api/middleware/admin"
	GlobalMiddleware "backend_capstone/api/middleware/global"
	UserMiddleware "backend_capstone/api/middleware/user"
	"backend_capstone/api/payment"
	"backend_capstone/api/product"
	"backend_capstone/api/productbrand"
	"backend_capstone/api/productcategory"
	"backend_capstone/api/transaction"
	"backend_capstone/api/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	Payment         *payment.Controller
	ProductCategory *productcategory.Controller
	ProductBrand    *productbrand.Controller
	Product         *product.Controller
	// PaymentMethod   *paymentmethod.Controller
	// PaymentVendor   *paymentvendor.Controller
	User        *user.Controller
	Transaction *transaction.Controller

	MiddlewareAdminJWT AdminMiddleware.JwtService
	MiddlewareUserJWT  UserMiddleware.JwtService
	MiddlewareJWT      GlobalMiddleware.JwtService
}

func RegistrationPath(e *echo.Echo, controller Controller) {
	// API v1 basepath
	e.GET("/v1", func(c echo.Context) error {
		return c.String(http.StatusOK, "Payzone API v1.0.0 Basepath")
	})

	e.POST("/v1/user_register", controller.User.Create)
	e.POST("/v1/admin_register", controller.User.CreateAdmin)

	authV1 := e.Group("/v1/auth")
	authV1.POST("", controller.User.AuthUser)

	transactionV1 := e.Group("/v1/transactions")
	transactionV1.Use(controller.MiddlewareAdminJWT.JwtAdminMiddleware())
	transactionV1.GET("/:id", controller.Transaction.GetById)
	// paymentV1.POST("/:method/:vendor", controller.Payment.Create)
	transactionV1.PUT("/:id", controller.Transaction.Modify)

	categoryV1 := e.Group("v1/product_categories")
	categoryV1.Use(controller.MiddlewareAdminJWT.JwtAdminMiddleware())
	categoryV1.POST("", controller.ProductCategory.Create)
	categoryV1.GET("", controller.ProductCategory.GetAll)
	categoryV1.GET("/:id", controller.ProductCategory.GetById)
	categoryV1.PUT("/:id", controller.ProductCategory.Modify)
	categoryV1.DELETE("/:id", controller.ProductCategory.Remove)

	brandV1 := e.Group("v1/product_brands")
	brandV1.Use(controller.MiddlewareAdminJWT.JwtAdminMiddleware())
	brandV1.POST("", controller.ProductBrand.Create)
	brandV1.GET("", controller.ProductBrand.GetAll)
	brandV1.GET("/:id", controller.ProductBrand.GetById)
	brandV1.PUT("/:id", controller.ProductBrand.Modify)
	brandV1.DELETE("/:id", controller.ProductBrand.Remove)
	brandV1.POST("/:id/categories/:category_id", controller.ProductBrand.AddBrandCategory)
	brandV1.DELETE("/:id/categories/:category_id", controller.ProductBrand.RemoveBrandCategory)

	productV1 := e.Group("v1/products")
	productV1.Use(controller.MiddlewareAdminJWT.JwtAdminMiddleware())
	productV1.POST("", controller.Product.Create)
	productV1.GET("", controller.Product.GetAll)
	productV1.GET("/:id", controller.Product.GetById)
	productV1.PUT("/:id", controller.Product.Modify)
	productV1.DELETE("/:id", controller.Product.Remove)

	userV1 := e.Group("v1/users")
	userV1.Use(controller.MiddlewareJWT.JwtMiddleware())
	userV1.POST("/transactions", controller.Transaction.Create)
	userV1.GET("/:id/transactions", controller.Transaction.UsersGetAll)
	userV1.GET("/:id", controller.User.GetSingleData)
	userV1.PUT("/:id", controller.User.UpdateUserData)
	userV1.DELETE("/:id", controller.User.DeleteData)
	// e.GET("/user", controller.User.GetAllData, middleware.IsAuthenticated)

	clientV1 := e.Group("v1/clients")
	clientV1.GET("/products", controller.Product.ClientGetAll)
	clientV1.GET("/products/:slug", controller.Product.ClientGetAllBySlug)
	clientV1.GET("/products/categories", controller.ProductCategory.GetAll)
	clientV1.GET("/payments", controller.Payment.GetAll)

	tokenV1 := e.Group("v1/tokens")
	tokenV1.Use(controller.MiddlewareJWT.JwtMiddleware())
	tokenV1.GET("", controller.User.ParseToken)

	midtransV1 := e.Group("v1/midtrans")
	midtransV1.POST("/bababoe", controller.Transaction.Modify)
}
