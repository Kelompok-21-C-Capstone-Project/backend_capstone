package product

import (
	"backend_capstone/api/product/request"
	"backend_capstone/api/product/response"
	productUseCase "backend_capstone/services/product"
	_ "backend_capstone/services/product/dto"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	service productUseCase.Service
}

func NewController(service productUseCase.Service) *Controller {
	return &Controller{
		service: service,
	}
}

// Create godoc
// @Summary Create product
// @Description  Create new product product
// @Tags         products
// @Accept       json
// @Produce      json
// @Param Payload body request.CreateProductRequest true "Payload format" SchemaExample(request.CreateProductRequest)
// @Success      201  {object}  models.Product
// @Failure      400  {object}  response.BasicProductResponse
// @Failure      403  {object}  response.BasicProductResponse
// @Failure      500  {object}  response.BasicProductResponse
// @Security ApiKeyAuth
// @Router       /v1/products [post]
func (controller *Controller) Create(c echo.Context) (err error) {
	log.Print("enter controller.product.Create")
	createProductReq := new(request.CreateProductRequest)
	if err := c.Bind(createProductReq); err != nil {
		return c.JSON(http.StatusBadRequest, response.BasicProductResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	data, err := controller.service.Create(*createProductReq.DtoReq())
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BasicProductResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, data)
}

// GetAll godoc
// @Summary Get product
// @Description  Get product data from database
// @Tags         products
// @Produce      json
// @Success      200  {array}  models.ProductBrand
// @Failure      400  {object}  response.BasicProductResponse
// @Failure      403  {object}  response.BasicProductResponse
// @Failure      500  {object}  response.BasicProductResponse
// @Security ApiKeyAuth
// @Router       /v1/products [get]
func (controller *Controller) GetAll(c echo.Context) (err error) {
	log.Print("enter controller.product.GetAll")
	datas, err := controller.service.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BasicProductResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, datas)
}

// ClientGetAll godoc
// @Summary Get products by all category for frontned
// @Description  Get all products by all category frontned from database
// @Tags         clients
// @Produce      json
// @Success      200  {array}  dto.ProductCategory
// @Failure      400  {object}  response.BasicProductResponse
// @Failure      403  {object}  response.BasicProductResponse
// @Failure      500  {object}  response.BasicProductResponse
// @Router       /v1/clients/products [get]
func (controller *Controller) ClientGetAll(c echo.Context) (err error) {
	log.Print("enter controller.product.GetAll")
	datas, err := controller.service.ClientGetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BasicProductResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, datas)
}

// ClientGetAllBySlug godoc
// @Summary Get product by specific category for frontend
// @Description  Get product by specific category for frontend from database
// @Tags         clients
// @Produce      json
// @Success      200  {object}  dto.ProductCategory
// @Failure      400  {object}  response.BasicProductResponse
// @Failure      403  {object}  response.BasicProductResponse
// @Failure      500  {object}  response.BasicProductResponse
// @Router       /v1/clients/products/:slug [get]
func (controller *Controller) ClientGetAllBySlug(c echo.Context) (err error) {
	log.Print("enter controller.product.GetAll")
	slug := c.Param("slug")
	datas, err := controller.service.ClientGetAllBySlug(slug)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BasicProductResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, datas)
}

// GetById godoc
// @Summary Get product
// @Description  Get product product by id
// @Tags         products
// @Produce      json
// @Param id   path  string  true  "Product ID" minLength:"32"
// @Success      200  {object}  models.Product
// @Failure      400  {object}  response.BasicProductResponse
// @Failure      403  {object}  response.BasicProductResponse
// @Failure      500  {object}  response.BasicProductResponse
// @Security ApiKeyAuth
// @Router       /v1/products/{id} [get]
func (controller *Controller) GetById(c echo.Context) (err error) {
	log.Print("enter controller.product.GetById")
	id := c.Param("id")
	data, err := controller.service.GetById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BasicProductResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, data)
}

// Modify godoc
// @Summary Update product
// @Description  Update product data
// @Tags         products
// @Accept       json
// @Produce      json
// @Param id   path  string  true  "Brand ID" minLength:"32"
// @Param Payload body request.UpdateProductRequest true "Payload format" SchemaExample(request.UpdateProductRequest)
// @Success      200  {object}  models.Product
// @Failure      400  {object}  response.BasicProductResponse
// @Failure      403  {object}  response.BasicProductResponse
// @Failure      500  {object}  response.BasicProductResponse
// @Security ApiKeyAuth
// @Router       /v1/products/{id} [put]
func (controller *Controller) Modify(c echo.Context) (err error) {
	log.Print("enter controller.product.Modify")
	id := c.Param("id")
	updateProductCategoryReq := new(request.UpdateProductRequest)
	if err = c.Bind(updateProductCategoryReq); err != nil {
		return c.JSON(http.StatusBadRequest, response.BasicProductResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	data, err := controller.service.Modify(id, *updateProductCategoryReq.DtoReq())
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BasicProductResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, data)
}

// Remove godoc
// @Summary Delete product data by id
// @Description  Delete product data from database
// @Tags         products
// @Produce      json
// @Param id   path  string  true  "Brand ID" minLength:"32"
// @Success      200  {object}  response.BasicProductSuccessResponse
// @Failure      400  {object}  response.BasicProductResponse
// @Failure      403  {object}  response.BasicProductResponse
// @Failure      500  {object}  response.BasicProductResponse
// @Security ApiKeyAuth
// @Router       /v1/products/{id} [delete]
func (controller *Controller) Remove(c echo.Context) (err error) {
	log.Print("enter controller.product.Remove")
	id := c.Param("id")
	if err = controller.service.Remove(id); err != nil {
		return c.JSON(http.StatusBadRequest, response.BasicProductResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusAccepted, response.BasicProductSuccessResponse{
		Status: "success",
	})
}
