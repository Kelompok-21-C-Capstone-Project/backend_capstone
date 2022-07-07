package transaction

import (
	"backend_capstone/api/transaction/request"
	"backend_capstone/api/transaction/response"
	transactionUseCase "backend_capstone/services/transaction"
	"errors"
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
// @Summary User create transaction
// @Description  Create new transaction for users
// @Tags         users
// @Accept       json
// @Produce      json
// @Param Payload body request.CreateTransactionRequest true "Payload format" SchemaExample(request.CreateTransactionRequest)
// @Success      201  {object}  dto.BillClient
// @Failure      400  {object}  response.BasicTransactionResponse
// @Failure      403  {object}  response.BasicTransactionResponse
// @Failure      500  {object}  response.BasicTransactionResponse
// @Security ApiKeyAuth
// @Router       /v1/users/transactions [post]
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

// GetAll GoDoc
// @Summary Get transaction
// @Description  Get transaction data
// @Tags         transactions
// @Produce      json
// @Success      200  {array}  dto.BillClient
// @Failure      400  {object}  response.BasicTransactionResponse
// @Failure      403  {object}  response.BasicTransactionResponse
// @Failure      500  {object}  response.BasicTransactionResponse
// @Security ApiKeyAuth
// @Router       /v1/transactions [get]
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

// UsersGetAll godoc
// @Summary Get all transaction from specific user
// @Description  Get all transaction from specific user
// @Param id   path  string  true  "user ID" minLength:"32"
// @Tags         users
// @Produce      json
// @Success      200  {array}  dto.ClientTransactionsResponse
// @Failure      400  {object}  response.BasicTransactionResponse
// @Failure      403  {object}  response.BasicTransactionResponse
// @Failure      500  {object}  response.BasicTransactionResponse
// @Security ApiKeyAuth
// @Router       /v1/users/:id/transactions [get]
func (controller *Controller) UsersGetAll(c echo.Context) (err error) {
	log.Print("enter controller.transaction.UsersGetAll")
	id := c.Param("id")
	// page := c.QueryParam("pages")
	// limit := c.QueryParam("limits")
	if id != c.Get("payload").(string) {
		err = errors.New("Tidak berizin")
		return
	}
	datas, err := controller.service.UsersGetAll(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BasicTransactionResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, datas)
}

// UsersGetById godoc
// @Summary Get all transaction from specific user
// @Description  Get all transaction from specific user
// @Param id   path  string  true  "user ID" minLength:"32"
// @Param transaction_id   path  string  true  "transaction ID" minLength:"32"
// @Tags         users
// @Produce      json
// @Success      200  {array}  dto.ClientTransactionsResponse
// @Failure      400  {object}  response.BasicTransactionResponse
// @Failure      403  {object}  response.BasicTransactionResponse
// @Failure      500  {object}  response.BasicTransactionResponse
// @Security ApiKeyAuth
// @Router       /v1/users/:id/transactions/:transaction_id [get]
func (controller *Controller) UsersGetById(c echo.Context) (err error) {
	log.Print("enter controller.transaction.UsersGetAll")
	id := c.Param("id")
	tid := c.Param("transaction_id")
	if id != c.Get("payload").(string) {
		err = errors.New("Tidak berizin")
		return
	}
	data, err := controller.service.UsersGetById(id, tid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BasicTransactionResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, data)
}

// GetById Godoc
// @Summary Get transaction
// @Description  Get transaction transaction by id
// @Tags         transactions
// @Produce      json
// @Param id   path  string  true  "transaction ID" minLength:"32"
// @Success      200  {object}  dto.BillClient
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

// GetBill Godoc
// @Summary Get transactions bill
// @Description  Get transaction transactions nill by id & user id
// @Tags         users
// @Produce      json
// @Param id   path  string  true  "transaction ID" minLength:"32"
// @Param transaction_id   path  string  true  "transaction ID" minLength:"32"
// @Success      200  {object}  dto.BillClient
// @Failure      400  {object}  response.BasicTransactionResponse
// @Failure      403  {object}  response.BasicTransactionResponse
// @Failure      500  {object}  response.BasicTransactionResponse
// @Security ApiKeyAuth
// @Router       /v1/users/{id}/transactions/{transaction_id}/bills [get]
func (controller *Controller) GetBill(c echo.Context) (err error) {
	id := c.Param("id")
	tid := c.Param("transaction_id")
	data, err := controller.service.GetBill(id, tid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BasicTransactionResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, data)
}
func (controller *Controller) MidtransAfterPayment(c echo.Context) (err error) {
	log.Print(c.Request().Body)
	reqMidtrans := new(request.MidtransReq)
	if err := c.Bind(reqMidtrans); err != nil {
		// return c.JSON(http.StatusBadRequest, response.BasicTransactionResponse{
		// 	Status:  "fail",
		// 	Message: err.Error(),
		// })
		log.Print(err)
	}
	if err := controller.service.MidtransAfterPayment(reqMidtrans.DtoReq()); err != nil {
		// return c.JSON(http.StatusBadRequest, response.BasicTransactionResponse{
		// 	Status:  "fail",
		// 	Message: err.Error(),
		// })
		log.Print(err)
	}
	return
}
