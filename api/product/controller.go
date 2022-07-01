package product

import (
	"backend_capstone/api/product/request"
	"backend_capstone/api/product/response"
	"backend_capstone/models"
	productUseCase "backend_capstone/services/product"
	"log"
	"net/http"
	"strconv"

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
func (controller *Controller) Remove(c echo.Context) (err error) {
	log.Print("enter controller.product.Remove")
	id := c.Param("id")
	if err = controller.service.Remove(id); err != nil {
		return c.JSON(http.StatusBadRequest, response.BasicProductResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusAccepted, response.BasicProductResponse{
		Status: "success",
	})
}

//> fungsi get (read) koreksi
func GetAllProduct(c echo.Context) error {
	var product []models.Product
	var error error
	if error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": error.Error()})
	}
	return c.JSON(http.StatusOK, product)
}

//> fungsi post(create) koreksi
func CreateProduct(c echo.Context) error {
	var input []models.Product
	err := c.Bind(&input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, input)
}

//> fungsi delete msh butuh koreksi dibagian return
func DeleteProduct(c echo.Context) error {
	var product map[int]*models.Product
	id, _ := strconv.Atoi(c.Param("id"))
	delete(product, id)
	return c.NoContent(http.StatusNoContent)
}

//>fungsi put(update) msh butuh koreksi
func UpdateProduct(c echo.Context) error {
	var product map[int]*models.Product
	n := new(models.Product)
	if err := c.Bind(n); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	product[id].Name = n.Name
	return c.JSON(http.StatusOK, product[id])
}
