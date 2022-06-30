package paymentmethod

import (
	"backend_capstone/api/paymentmethod/request"
	"backend_capstone/api/paymentmethod/response"
	paymentmethodUseCase "backend_capstone/services/paymentmethod"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	service paymentmethodUseCase.Service
}

func NewController(service paymentmethodUseCase.Service) *Controller {
	return &Controller{
		service: service,
	}
}

func (controller *Controller) Create(c echo.Context) (err error) {
	log.Print("enter controller.paymentmethod.Create")
	createPaymentMethodReq := new(request.CreateUpdateMethodRequest)
	if err := c.Bind(createPaymentMethodReq); err != nil {
		return c.JSON(http.StatusBadRequest, response.BasicMethodResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	data, err := controller.service.Create(*createPaymentMethodReq.DtoReq())
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BasicMethodResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, data)
}
func (controller *Controller) GetAll(c echo.Context) (err error) {
	log.Print("enter controller.paymentmethod.GetAll")
	datas, err := controller.service.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BasicMethodResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, datas)
}
func (controller *Controller) GetById(c echo.Context) (err error) {
	log.Print("enter controller.paymentmethod.GetById")
	id := c.Param("id")
	data, err := controller.service.GetById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BasicMethodResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, data)
}
func (controller *Controller) Modify(c echo.Context) (err error) {
	log.Print("enter controller.paymentmethod.Modify")
	id := c.Param("id")
	updatePaymentMethodReq := new(request.CreateUpdateMethodRequest)
	if err = c.Bind(updatePaymentMethodReq); err != nil {
		return c.JSON(http.StatusBadRequest, response.BasicMethodResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	data, err := controller.service.Modify(id, *updatePaymentMethodReq.DtoReq())
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BasicMethodResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, data)
}
func (controller *Controller) Remove(c echo.Context) (err error) {
	log.Print("enter controller.paymentmethod.Remove")
	id := c.Param("id")
	if err = controller.service.Remove(id); err != nil {
		return c.JSON(http.StatusBadRequest, response.BasicMethodResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusAccepted, response.BasicMethodResponse{
		Status: "success",
	})
}
