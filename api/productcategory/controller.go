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

// Create godoc
// @Summary Create category
// @Description  Create new product category
// @Tags         product_categories
// @Accept       json
// @Produce      json
// @Param Payload body request.CreateCategoryRequest true "Payload format" SchemaExample(request.CreateCategoryRequest)
// @Success      201  {object}  models.ProductCategory
// @Failure      400  {object}  response.BasicCategoryResponse
// @Failure      403  {object}  response.BasicCategoryResponse
// @Failure      500  {object}  response.BasicCategoryResponse
// @Security ApiKeyAuth
// @Router       /v1/product_categories [post]
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

// GetById godoc
// @Summary Get category
// @Description  Get product category by id
// @Tags         product_categories
// @Produce      json
// @Param id   path  string  true  "Category ID" minLength:"32"
// @Success      200  {object}  models.ProductCategory
// @Failure      400  {object}  response.BasicCategoryResponse
// @Failure      403  {object}  response.BasicCategoryResponse
// @Failure      500  {object}  response.BasicCategoryResponse
// @Security ApiKeyAuth
// @Router       /v1/product_categories/{id} [get]
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

// Modify godoc
// @Summary Update category
// @Description  Update category data
// @Tags         product_categories
// @Accept       json
// @Produce      json
// @Param id   path  string  true  "Category ID" minLength:"32"
// @Param Payload body request.UpdateCategoryRequest true "Payload format" SchemaExample(request.UpdateCategoryRequest)
// @Success      200  {object}  models.ProductCategory
// @Failure      400  {object}  response.BasicCategoryResponse
// @Failure      403  {object}  response.BasicCategoryResponse
// @Failure      500  {object}  response.BasicCategoryResponse
// @Security ApiKeyAuth
// @Router       /v1/product_categories/{id} [put]
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

// Remove godoc
// @Summary Delete category data by id
// @Description  Delete category data from database
// @Tags         product_categories
// @Produce      json
// @Param id   path  string  true  "Category ID" minLength:"32"
// @Success      200  {object}  response.BasicCategorySuccessResponse
// @Failure      400  {object}  response.BasicCategoryResponse
// @Failure      403  {object}  response.BasicCategoryResponse
// @Failure      500  {object}  response.BasicCategoryResponse
// @Security ApiKeyAuth
// @Router       /v1/product_categories/{id} [delete]
func (controller *Controller) Remove(c echo.Context) (err error) {
	log.Print("enter controller.productcategory.Remove")
	id := c.Param("id")
	if err = controller.service.Remove(id); err != nil {
		return c.JSON(http.StatusBadRequest, response.BasicCategoryResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusAccepted, response.BasicCategorySuccessResponse{
		Status: "success",
	})
}
