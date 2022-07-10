package productbrand

import (
	"backend_capstone/api/productbrand/request"
	"backend_capstone/api/productbrand/response"
	productbrandUseCase "backend_capstone/services/productbrand"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	service productbrandUseCase.Service
}

func NewController(service productbrandUseCase.Service) *Controller {
	return &Controller{
		service: service,
	}
}

// Create godoc
// @Summary Create brand
// @Description  Create new product brand
// @Tags         product_brands
// @Accept       json
// @Produce      json
// @Param Payload body request.CreateBrandRequest true "Payload format" SchemaExample(request.CreateBrandRequest)
// @Success      201  {object}  models.ProductBrand
// @Failure      400  {object}  response.BasicBrandResponse
// @Failure      403  {object}  response.BasicBrandResponse
// @Failure      500  {object}  response.BasicBrandResponse
// @Security ApiKeyAuth
// @Router       /v1/product_brands [post]
func (controller *Controller) Create(c echo.Context) (err error) {
	log.Print("enter controller.productbrand.Create")
	createProductBrandReq := new(request.CreateBrandRequest)
	if err := c.Bind(createProductBrandReq); err != nil {
		return c.JSON(http.StatusBadRequest, response.BasicBrandResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	data, err := controller.service.Create(*createProductBrandReq.DtoReq())
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BasicBrandResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, data)
}

// GetAll godoc
// @Summary Get brand
// @Description  Get product brand by id
// @Param query   query  string  false  "search data by query"
// @Param page   query  string  false  "search data by page"
// @Param page_size   query  string  false  "search data by page size"
// @Tags         product_brands
// @Produce      json
// @Success      200  {array}  dto.ResponseBodyProductBrand
// @Failure      400  {object}  response.BasicBrandResponse
// @Failure      403  {object}  response.BasicBrandResponse
// @Failure      500  {object}  response.BasicBrandResponse
// @Security ApiKeyAuth
// @Router       /v1/product_brands [get]
func (controller *Controller) GetAll(c echo.Context) (err error) {
	log.Print("enter controller.productbrand.GetAll")
	query := c.QueryParam("query")
	page := c.QueryParam("page")
	pageSize := c.QueryParam("page_size")
	datas, err := controller.service.GetAll(query, page, pageSize)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BasicBrandResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, datas)
}

// GetById godoc
// @Summary Get brand
// @Description  Get product brand by id
// @Tags         product_brands
// @Produce      json
// @Param id   path  string  true  "Brand ID" minLength:"32"
// @Success      200  {object}  models.ProductBrand
// @Failure      400  {object}  response.BasicBrandResponse
// @Failure      403  {object}  response.BasicBrandResponse
// @Failure      500  {object}  response.BasicBrandResponse
// @Security ApiKeyAuth
// @Router       /v1/product_brands/{id} [get]
func (controller *Controller) GetById(c echo.Context) (err error) {
	log.Print("enter controller.productbrand.GetById")
	id := c.Param("id")
	data, err := controller.service.GetById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BasicBrandResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, data)
}

// Modify godoc
// @Summary Update brand
// @Description  Update brand data
// @Tags         product_brands
// @Accept       json
// @Produce      json
// @Param id   path  string  true  "Brand ID" minLength:"32"
// @Param Payload body request.UpdateBrandRequest true "Payload format" SchemaExample(request.UpdateBrandRequest)
// @Success      200  {object}  models.ProductBrand
// @Failure      400  {object}  response.BasicBrandResponse
// @Failure      403  {object}  response.BasicBrandResponse
// @Failure      500  {object}  response.BasicBrandResponse
// @Security ApiKeyAuth
// @Router       /v1/product_brands/{id} [put]
func (controller *Controller) Modify(c echo.Context) (err error) {
	log.Print("enter controller.productbrand.Modify")
	id := c.Param("id")
	updateProductCategoryReq := new(request.UpdateBrandRequest)
	if err = c.Bind(updateProductCategoryReq); err != nil {
		return c.JSON(http.StatusBadRequest, response.BasicBrandResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	data, err := controller.service.Modify(id, *updateProductCategoryReq.DtoReq())
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BasicBrandResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, data)
}

// Remove godoc
// @Summary Delete brand data by id
// @Description  Delete brand data from database
// @Tags         product_brands
// @Produce      json
// @Param id   path  string  true  "Brand ID" minLength:"32"
// @Success      200  {object}  response.BasicBrandSuccessResponse
// @Failure      400  {object}  response.BasicBrandResponse
// @Failure      403  {object}  response.BasicBrandResponse
// @Failure      500  {object}  response.BasicBrandResponse
// @Security ApiKeyAuth
// @Router       /v1/product_brands/{id} [delete]
func (controller *Controller) Remove(c echo.Context) (err error) {
	log.Print("enter controller.productbrand.Remove")
	id := c.Param("id")
	if err = controller.service.Remove(id); err != nil {
		return c.JSON(http.StatusBadRequest, response.BasicBrandResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusAccepted, response.BasicBrandSuccessResponse{
		Status: "success",
	})
}

// AddBrandCategory godoc
// @Summary Add category to brand
// @Description  Add category to brand
// @Tags         product_brands
// @Produce      json
// @Param id   path  string  true  "Brand ID" minLength:"32"
// @Param category_id   path  string  true  "Category ID" minLength:"32"
// @Success      200  {object}  response.BasicBrandSuccessResponse
// @Failure      400  {object}  response.BasicBrandResponse
// @Failure      403  {object}  response.BasicBrandResponse
// @Failure      500  {object}  response.BasicBrandResponse
// @Security ApiKeyAuth
// @Router       /v1/product_brands/{id}/categories/{category_id} [post]
func (controller *Controller) AddBrandCategory(c echo.Context) (err error) {
	brandId := c.Param("id")
	categoryId := c.Param("category_id")
	_, err = controller.service.AddBrandCategory(brandId, categoryId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BasicBrandResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusAccepted, response.BasicBrandSuccessResponse{
		Status: "success",
	})
}

// RemoveBrandCategory godoc
// @Summary Remove category from brand
// @Description  Remove category from brand
// @Tags         product_brands
// @Produce      json
// @Param id   path  string  true  "Brand ID" minLength:"32"
// @Param category_id   path  string  true  "Category ID" minLength:"32"
// @Success      200  {object}  response.BasicBrandSuccessResponse
// @Failure      400  {object}  response.BasicBrandResponse
// @Failure      403  {object}  response.BasicBrandResponse
// @Failure      500  {object}  response.BasicBrandResponse
// @Security ApiKeyAuth
// @Router       /v1/product_brands/{id}/categories/{category_id} [delete]
func (controller *Controller) RemoveBrandCategory(c echo.Context) (err error) {
	brandId := c.Param("id")
	categoryId := c.Param("category_id")
	err = controller.service.RemoveBrandCategory(brandId, categoryId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BasicBrandResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusAccepted, response.BasicBrandSuccessResponse{
		Status: "success",
	})
}
