package main

import (
	"backend_capstone/api"
	"backend_capstone/configs"
	"backend_capstone/modules"
	"backend_capstone/utils"
	"backend_capstone/utils/gormdriver"
	"backend_capstone/utils/midtransdriver"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func main() {

	// Initialize echo
	e := echo.New()

	// Get server config variable
	config := configs.GetConfig()

	// Infrastructure driver adapters initialization
	// Database Adapters
	dbCon := utils.NewDatabaseConnection(config)

	// gorm migration
	gormMigrationService := gormdriver.NewGormMigrationService(dbCon)
	gormMigrationService.GormMigrate()

	defer dbCon.CloseConnection()

	// Midtrans Adapters
	midtransDriver := midtransdriver.NewMidtransService(config)

	// Interface driving adapters initialization
	// API Adapters
	controllers := modules.RegisterModules(dbCon, midtransDriver, config)
	api.RegistrationPath(e, controllers)

	// Swagger api documentation
	swaggerHandler := echoSwagger.WrapHandler

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Payzone API")
	})
	e.GET("/swagger/*", swaggerHandler)

	go func() {
		address := fmt.Sprintf("0.0.0.0:%d", config.App.PORT)
		if err := e.Start(address); err != nil {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal)
	<-quit
}
