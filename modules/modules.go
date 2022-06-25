package modules

import (
	"backend_capstone/api"
	"backend_capstone/configs"
	"backend_capstone/models"
	"backend_capstone/utils"
)

func RegisterModules(dbCon *utils.DatabaseConnection, configs *configs.AppConfig) api.Controller {
	// gorm migration
	gormMigrationService := models.NewGormMigrationService(dbCon)
	gormMigrationService.GormMigrate()

	controllers := api.Controller{}

	return controllers
}
