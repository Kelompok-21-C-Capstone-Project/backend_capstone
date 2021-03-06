package payment

import (
	"backend_capstone/api/payment/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controller struct {
}

var (
	paymentMethods = []response.PaymentMethods{
		{
			Id:   1,
			Type: "Virtual Account",
			Services: []response.PaymentService{
				{
					Id:    1,
					Label: "BCA Virtual Account",
					Icon:  "BCA",
				},
				{
					Id:    2,
					Label: "BNI Virtual Account",
					Icon:  "BNI",
				},
				{
					Id:    3,
					Label: "Mandiri Virtual Account",
					Icon:  "Mandiri",
				},
				{
					Id:    4,
					Label: "BRI Virtual Account",
					Icon:  "BRI",
				},
				{
					Id:    5,
					Label: "Permata Virtual Account",
					Icon:  "Permata",
				},
				// {
				// 	Id:    4,
				// 	Label: "CIMB Virtual Account",
				// 	Icon:  "CIMB",
				// },
				// {
				// 	Id:    6,
				// 	Label: "Maybank Virtual Account",
				// 	Icon:  "Maybank",
				// },
				// {
				// 	Id:    8,
				// 	Label: "Mega Virtual Account",
				// 	Icon:  "Mega",
				// },
			},
		},
		{
			Id:   2,
			Type: "Gopay",
			Services: []response.PaymentService{
				{
					Id:    1,
					Label: "Gopay",
					Icon:  "Gopay",
				},
			},
		},
	}
)

func NewController() *Controller {
	return &Controller{}
}

// GetAll godoc
// @Summary Get all payment
// @Description  Get all payment methods and services
// @Tags         clients
// @Produce      json
// @Success      200  {object}  response.PaymentMethods
// @Router       /v1/clients/payments [get]
func (controller *Controller) GetAll(c echo.Context) (err error) {
	return c.JSON(http.StatusOK, paymentMethods)
}
