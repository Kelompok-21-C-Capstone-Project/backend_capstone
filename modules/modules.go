package modules

import (
	"backend_capstone/api"
	brandController "backend_capstone/api/productbrand"
	categoryController "backend_capstone/api/productcategory"
	"backend_capstone/configs"
	brandRepo "backend_capstone/repository/productbrand"
	categoryRepo "backend_capstone/repository/productcategory"
	brandService "backend_capstone/services/productbrand"
	categoryService "backend_capstone/services/productcategory"
	"backend_capstone/utils"
	"backend_capstone/utils/midtransdriver"
)

func RegisterModules(dbCon *utils.DatabaseConnection, midtransDriver *midtransdriver.MidtransDriver, configs *configs.AppConfig) api.Controller {

	categoryPermitRepository := categoryRepo.RepositoryFactory(dbCon)
	categoryPermitService := categoryService.NewService(categoryPermitRepository)
	categoryPermitController := categoryController.NewController(categoryPermitService)

	brandPermitRepository := brandRepo.RepositoryFactory(dbCon)
	brandPermitService := brandService.NewService(brandPermitRepository)
	brandPermitController := brandController.NewController(brandPermitService)

	controllers := api.Controller{
		ProductCategory: categoryPermitController,
		ProductBrand:    brandPermitController,
	}

	return controllers
}
