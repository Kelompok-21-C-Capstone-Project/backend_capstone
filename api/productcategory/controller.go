package productcategory

import (
	"backend_capstone/api/productcategory/request"
	"backend_capstone/api/productcategory/response"
	productcategoryUseCase "backend_capstone/services/productcategory"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	service productcategoryUseCase.Service
}

func NewController(service productcategoryUseCase.Service) *Controller {
	return &Controller{
		service: service,
	}
}

func (controller *Controller) Create(c echo.Context) (err error) {
	log.Print("enter controller.productcategory.Create")
	createProductCategoryReq := new(request.CreateCategoryRequest)
	if err := c.Bind(createProductCategoryReq); err != nil {
		return c.JSON(http.StatusBadRequest, response.BasicCategoryResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	data, err := controller.service.Create(*createProductCategoryReq.DtoReq())
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BasicCategoryResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, data)
}
func (controller *Controller) GetAll(c echo.Context) (err error) {
	log.Print("enter controller.productcategory.GetAll")
	datas, err := controller.service.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BasicCategoryResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, datas)
}
func (controller *Controller) GetById(c echo.Context) (err error) {
	log.Print("enter controller.productcategory.GetById")
	id := c.Param("id")
	data, err := controller.service.GetById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BasicCategoryResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, data)
}
func (controller *Controller) Modify(c echo.Context) (err error) {
	log.Print("enter controller.productcategory.Modify")
	id := c.Param("id")
	updateProductCategoryReq := new(request.UpdateCategoryRequest)
	if err = c.Bind(updateProductCategoryReq); err != nil {
		return c.JSON(http.StatusBadRequest, response.BasicCategoryResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	data, err := controller.service.Modify(id, *updateProductCategoryReq.DtoReq())
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BasicCategoryResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, data)
}
func (controller *Controller) Remove(c echo.Context) (err error) {
	log.Print("enter controller.productcategory.Remove")
	id := c.Param("id")
	if err = controller.service.Remove(id); err != nil {
		return c.JSON(http.StatusBadRequest, response.BasicCategoryResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusAccepted, response.BasicCategoryResponse{
		Status: "success",
	})
}
