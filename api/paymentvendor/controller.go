package paymentvendor

import (
	"backend_capstone/api/paymentvendor/request"
	"backend_capstone/api/paymentvendor/response"
	paymentvendorUseCase "backend_capstone/services/paymentvendor"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	service paymentvendorUseCase.Service
}

func NewController(service paymentvendorUseCase.Service) *Controller {
	return &Controller{
		service: service,
	}
}

func (controller *Controller) Create(c echo.Context) (err error) {
	log.Print("enter controller.productbrand.Create")
	createProductBrandReq := new(request.CreateVendorRequest)
	if err := c.Bind(createProductBrandReq); err != nil {
		return c.JSON(http.StatusBadRequest, response.BasicVendorResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	data, err := controller.service.Create(*createProductBrandReq.DtoReq())
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BasicVendorResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, data)
}
func (controller *Controller) GetAll(c echo.Context) (err error) {
	log.Print("enter controller.productbrand.GetAll")
	datas, err := controller.service.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BasicVendorResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, datas)
}
func (controller *Controller) GetById(c echo.Context) (err error) {
	log.Print("enter controller.productbrand.GetById")
	id := c.Param("id")
	data, err := controller.service.GetById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BasicVendorResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, data)
}
func (controller *Controller) Modify(c echo.Context) (err error) {
	log.Print("enter controller.productbrand.Modify")
	id := c.Param("id")
	updateProductCategoryReq := new(request.UpdateVendorRequest)
	if err = c.Bind(updateProductCategoryReq); err != nil {
		return c.JSON(http.StatusBadRequest, response.BasicVendorResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	data, err := controller.service.Modify(id, *updateProductCategoryReq.DtoReq())
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BasicVendorResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, data)
}
func (controller *Controller) Remove(c echo.Context) (err error) {
	log.Print("enter controller.productbrand.Remove")
	id := c.Param("id")
	if err = controller.service.Remove(id); err != nil {
		return c.JSON(http.StatusBadRequest, response.BasicVendorResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusAccepted, response.BasicVendorResponse{
		Status: "success",
	})
}
