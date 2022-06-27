package modules

import (
	"backend_capstone/api"
	paymentController "backend_capstone/api/payment"
	"backend_capstone/configs"
	paymentService "backend_capstone/services/payment"
	"backend_capstone/utils"
	"backend_capstone/utils/midtransdriver"
	"log"
)

func RegisterModules(dbCon *utils.DatabaseConnection, midtransDriver *midtransdriver.MidtransDriver, configs *configs.AppConfig) api.Controller {
	log.Print("Enter RegisterModules")
	paymentPermitService := paymentService.NewService(nil, midtransDriver)
	paymentV1PermitController := paymentController.NewController(paymentPermitService)

	controllers := api.Controller{
		PaymentV1Controller: paymentV1PermitController,
	}

	return controllers
}
