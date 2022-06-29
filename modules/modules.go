package modules

import (
	"backend_capstone/api"
	categoryController "backend_capstone/api/productcategory"
	"backend_capstone/configs"
	categoryRepo "backend_capstone/repository/productcategory"
	categoryService "backend_capstone/services/productcategory"
	"backend_capstone/utils"
	"backend_capstone/utils/midtransdriver"
)

func RegisterModules(dbCon *utils.DatabaseConnection, midtransDriver *midtransdriver.MidtransDriver, configs *configs.AppConfig) api.Controller {

	categoryPermitRepository := categoryRepo.RepositoryFactory(dbCon)
	categoryPermitService := categoryService.NewService(categoryPermitRepository)
	categoryPermitController := categoryController.NewController(categoryPermitService)

	controllers := api.Controller{
		ProductCategory: categoryPermitController,
	}

	return controllers
}
