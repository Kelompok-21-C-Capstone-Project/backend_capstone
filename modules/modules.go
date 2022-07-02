package modules

import (
	"backend_capstone/api"
	paymentController "backend_capstone/api/payment"
	methodController "backend_capstone/api/paymentmethod"
	vendorController "backend_capstone/api/paymentvendor"
	productController "backend_capstone/api/product"
	brandController "backend_capstone/api/productbrand"
	categoryController "backend_capstone/api/productcategory"
	"backend_capstone/api/user"
	"backend_capstone/configs"
	methodRepo "backend_capstone/repository/paymentmethod"
	vendorRepo "backend_capstone/repository/paymentvendor"
	productRepo "backend_capstone/repository/product"
	brandRepo "backend_capstone/repository/productbrand"
	categoryRepo "backend_capstone/repository/productcategory"
	userRepo "backend_capstone/repository/user"
	paymentService "backend_capstone/services/payment"
	methodService "backend_capstone/services/paymentmethod"
	vendorService "backend_capstone/services/paymentvendor"
	productService "backend_capstone/services/product"
	brandService "backend_capstone/services/productbrand"
	categoryService "backend_capstone/services/productcategory"
	userService "backend_capstone/services/user"
	"backend_capstone/utils"
	"backend_capstone/utils/midtransdriver"
	"backend_capstone/utils/security"

	"log"
)

func RegisterModules(dbCon *utils.DatabaseConnection, midtransDriver *midtransdriver.MidtransDriver, configs *configs.AppConfig) api.Controller {
	log.Print("Enter RegisterModules")

	passwordHashPermitUtils := security.NewPasswordHash()

	paymentPermitService := paymentService.NewService(nil, midtransDriver)
	paymentV1PermitController := paymentController.NewController(paymentPermitService)

	categoryPermitRepository := categoryRepo.RepositoryFactory(dbCon)
	categoryPermitService := categoryService.NewService(categoryPermitRepository)
	categoryPermitController := categoryController.NewController(categoryPermitService)

	brandPermitRepository := brandRepo.RepositoryFactory(dbCon)
	brandPermitService := brandService.NewService(brandPermitRepository)
	brandPermitController := brandController.NewController(brandPermitService)

	productPermitRepository := productRepo.RepositoryFactory(dbCon)
	productPermitService := productService.NewService(productPermitRepository)
	productPermitController := productController.NewController(productPermitService)

	methodPermitRepository := methodRepo.RepositoryFactory(dbCon)
	methodPermitService := methodService.NewService(methodPermitRepository)
	methodPermitController := methodController.NewController(methodPermitService)

	vendorPermitRepository := vendorRepo.RepositoryFactory(dbCon)
	vendorPermitService := vendorService.NewService(vendorPermitRepository)
	vendorPermitController := vendorController.NewController(vendorPermitService)

	userPermitRepository := userRepo.RepositoryFactory(dbCon)
	userPermitService := userService.NewService(userPermitRepository, passwordHashPermitUtils)
	userPermitController := user.NewController(userPermitService)

	controllers := api.Controller{
		ProductCategory: categoryPermitController,
		ProductBrand:    brandPermitController,
		Product:         productPermitController,
		PaymentMethod:   methodPermitController,
		PaymentVendor:   vendorPermitController,
		Payment:         paymentV1PermitController,
		User:            userPermitController,
	}

	return controllers
}
