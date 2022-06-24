package modules

import (
	"backend_capstone/api"
	"backend_capstone/configs"
	"backend_capstone/utils"
	"backend_capstone/utils/midtransdriver"
)

func RegisterModules(dbCon *utils.DatabaseConnection, midtransDriver *midtransdriver.MidtransDriver, configs *configs.AppConfig) api.Controller {
	controllers := api.Controller{}

	return controllers
}
