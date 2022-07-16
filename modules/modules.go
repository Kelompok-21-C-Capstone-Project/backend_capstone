package modules

import (
	"backend_capstone/api"
	AdminMiddleware "backend_capstone/api/middleware/admin"
	GlobalMiddleware "backend_capstone/api/middleware/global"
	UserMiddleware "backend_capstone/api/middleware/user"
	paymentController "backend_capstone/api/payment"
	productController "backend_capstone/api/product"
	brandController "backend_capstone/api/productbrand"
	categoryController "backend_capstone/api/productcategory"
	transactionController "backend_capstone/api/transaction"
	"backend_capstone/api/user"
	"backend_capstone/configs"
	productRepo "backend_capstone/repository/product"
	brandRepo "backend_capstone/repository/productbrand"
	categoryRepo "backend_capstone/repository/productcategory"
	transactionRepo "backend_capstone/repository/transaction"
	userRepo "backend_capstone/repository/user"
	paymentService "backend_capstone/services/payment"
	productService "backend_capstone/services/product"
	brandService "backend_capstone/services/productbrand"
	categoryService "backend_capstone/services/productcategory"
	transactionService "backend_capstone/services/transaction"
	userService "backend_capstone/services/user"
	"backend_capstone/utils"
	"backend_capstone/utils/mailjetdriver"
	"backend_capstone/utils/midtransdriver"
	"backend_capstone/utils/security"
	"log"
)

func RegisterModules(dbCon *utils.DatabaseConnection, configs *configs.AppConfig) api.Controller {
	log.Print("Enter RegisterModules")

	jwtPermitUtils := security.NewJwtService(configs.App.JWT)
	passwordHashPermitUtils := security.NewPasswordHash()
	midtransPermitUtils := midtransdriver.NewMidtransService(configs)
	mailjetPermitUtils := mailjetdriver.NewTransactionMailjetService(configs)

	transactionPermitRepository := transactionRepo.RepositoryFactory(dbCon)
	transactionPermitService := transactionService.NewService(transactionPermitRepository, midtransPermitUtils, mailjetPermitUtils)
	transactionPermitController := transactionController.NewController(transactionPermitService)

	paymentPermitService := paymentService.NewService(nil)
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

	// methodPermitRepository := methodRepo.RepositoryFactory(dbCon)
	// methodPermitService := methodService.NewService(methodPermitRepository)
	// methodPermitController := methodController.NewController(methodPermitService)

	// vendorPermitRepository := vendorRepo.RepositoryFactory(dbCon)
	// vendorPermitService := vendorService.NewService(vendorPermitRepository)
	// vendorPermitController := vendorController.NewController(vendorPermitService)

	userPermitRepository := userRepo.RepositoryFactory(dbCon)
	userPermitService := userService.NewService(userPermitRepository, passwordHashPermitUtils, jwtPermitUtils, configs.App.ADMIN_KEY)
	userPermitController := user.NewController(userPermitService)

	middlewarePermitUser := UserMiddleware.NewJwtUserMiddleware(configs.App.JWT)
	middlewarePermitAdmin := AdminMiddleware.NewJwtAdminMiddleware(configs.App.JWT)
	middlewarePermitGlobal := GlobalMiddleware.NewJwtMiddleware(configs.App.JWT)

	controllers := api.Controller{
		ProductCategory: categoryPermitController,
		ProductBrand:    brandPermitController,
		Product:         productPermitController,
		// PaymentMethod:      methodPermitController,
		// PaymentVendor:      vendorPermitController,
		Payment:            paymentV1PermitController,
		User:               userPermitController,
		MiddlewareUserJWT:  middlewarePermitUser,
		MiddlewareAdminJWT: middlewarePermitAdmin,
		Transaction:        transactionPermitController,
		MiddlewareJWT:      middlewarePermitGlobal,
	}

	return controllers
}
