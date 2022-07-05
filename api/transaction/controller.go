package transaction

import (
	"backend_capstone/api/transaction/request"
	"backend_capstone/api/transaction/response"
	transactionUseCase "backend_capstone/services/transaction"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	service transactionUseCase.Service
}

func NewController(service transactionUseCase.Service) *Controller {
	return &Controller{
		service: service,
	}
}

// Create godoc
// @Summary Create transaction
// @Description  Create new transaction for users
// @Tags         transactions
// @Accept       json
// @Produce      json
// @Param Payload body request.CreateTransactionRequest true "Payload format" SchemaExample(request.CreateTransactionRequest)
// @Success      201  {object}  dto.BillClient
// @Failure      400  {object}  response.BasicTransactionResponse
// @Failure      403  {object}  response.BasicTransactionResponse
// @Failure      500  {object}  response.BasicTransactionResponse
// @Security ApiKeyAuth
// @Router       /v1/transactions [post]
func (controller *Controller) Create(c echo.Context) (err error) {
	log.Print("enter controller.transaction.Create")
	payloadId := c.Get("payload").(string)
	createProductReq := new(request.CreateTransactionRequest)
	if err := c.Bind(createProductReq); err != nil {
		return c.JSON(http.StatusBadRequest, response.BasicTransactionResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	data, err := controller.service.Create(payloadId, createProductReq.DtoReq())
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BasicTransactionResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, data)
}

// @Summary Get transaction
// @Description  Get transaction transaction by id
// @Tags         users
// @Produce      json
// @Success      200  {array}  models.Transaction
// @Failure      400  {object}  response.BasicTransactionResponse
// @Failure      403  {object}  response.BasicTransactionResponse
// @Failure      500  {object}  response.BasicTransactionResponse
// @Security ApiKeyAuth
// @Router       /v1/users/transactions [post]
func (controller *Controller) GetAll(c echo.Context) (err error) {
	log.Print("enter controller.transaction.GetAll")
	datas, err := controller.service.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BasicTransactionResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, datas)
}

// @Summary Get transaction
// @Description  Get transaction transaction by id
// @Tags         transactions
// @Produce      json
// @Param id   path  string  true  "transaction ID" minLength:"32"
// @Success      200  {object}  models.Transaction
// @Failure      400  {object}  response.BasicTransactionResponse
// @Failure      403  {object}  response.BasicTransactionResponse
// @Failure      500  {object}  response.BasicTransactionResponse
// @Security ApiKeyAuth
// @Router       /v1/transactions/{id} [get]
func (controller *Controller) GetById(c echo.Context) (err error) {
	log.Print("enter controller.transaction.GetById")
	id := c.Param("id")
	data, err := controller.service.GetById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BasicTransactionResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, data)
}

// @Summary Update transaction
// @Description  Update transaction data
// @Tags         transactions
// @Accept       json
// @Produce      json
// @Param id   path  string  true  "Brand ID" minLength:"32"
// @Param Payload body request.UpdateProductRequest true "Payload format" SchemaExample(request.UpdateProductRequest)
// @Success      200  {object}  models.Transaction
// @Failure      400  {object}  response.BasicTransactionResponse
// @Failure      403  {object}  response.BasicTransactionResponse
// @Failure      500  {object}  response.BasicTransactionResponse
// @Security ApiKeyAuth
// @Router       /v1/transactions/{id} [put]
func (controller *Controller) Modify(c echo.Context) (err error) {
	// log.Print("enter controller.transaction.Modify")
	// id := c.Param("id")
	// updateProductCategoryReq := new(request.UpdateProductRequest)
	// if err = c.Bind(updateProductCategoryReq); err != nil {
	// 	return c.JSON(http.StatusBadRequest, response.BasicTransactionResponse{
	// 		Status:  "fail",
	// 		Message: err.Error(),
	// 	})
	// }
	// data, err := controller.service.Modify(id, *updateProductCategoryReq.DtoReq())
	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, response.BasicTransactionResponse{
	// 		Status:  "fail",
	// 		Message: err.Error(),
	// 	})
	// }
	data := new(interface{})
	if err := c.Bind(data); err != nil {
		return c.JSON(http.StatusBadRequest, response.BasicTransactionResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	log.Print(data)
	// return c.JSON(http.StatusOK, data)
	return
}

// @Summary Delete transaction data by id
// @Description  Delete transaction data from database
// @Tags         transactions
// @Produce      json
// @Param id   path  string  true  "Brand ID" minLength:"32"
// @Success      200  {object}  response.BasicProductSuccessResponse
// @Failure      400  {object}  response.BasicTransactionResponse
// @Failure      403  {object}  response.BasicTransactionResponse
// @Failure      500  {object}  response.BasicTransactionResponse
// @Security ApiKeyAuth
// @Router       /v1/transactions/{id} [delete]
func (controller *Controller) Remove(c echo.Context) (err error) {
	// log.Print("enter controller.transaction.Remove")
	// id := c.Param("id")
	// if err = controller.service.Remove(id); err != nil {
	// 	return c.JSON(http.StatusBadRequest, response.BasicTransactionResponse{
	// 		Status:  "fail",
	// 		Message: err.Error(),
	// 	})
	// }
	// return c.JSON(http.StatusAccepted, response.BasicTransactionSuccessResponse{
	// 	Status: "success",
	// })
	return
}
