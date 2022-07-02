package user

import (
	"backend_capstone/api/user/request"
	"backend_capstone/api/user/response"
	"backend_capstone/services/user"
	"log"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	service user.Service
}

func NewController(service user.Service) *Controller {
	return &Controller{
		service: service,
	}
}

func (ctrl *Controller) GetAllData(c echo.Context) (err error) {
	log.Print("enter ctrl.user.GetAllData")
	// user, err := utils.ParsingJWT(c)
	// else if user.Role != "admin" {
	// 	return c.JSON(200, echo.Map{
	// 		"error": "restricted (*only for admin)",
	// 	})
	// }
	users, err := ctrl.service.GetAll()
	if err != nil {
		return c.JSON(500, &response.BasicUserResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	return c.JSON(200, echo.Map{
		"data": users,
	})
}

func (ctrl *Controller) Create(c echo.Context) (err error) {
	log.Print("enter ctrl.user.Create")

	data := new(request.RegisterUserRequest)
	err = c.Bind(&data)
	if err != nil {
		return c.JSON(400, &response.BasicUserResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	user, err := ctrl.service.Create(data.DtoReq())
	if err != nil {
		return c.JSON(500, &response.BasicUserResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	return c.JSON(201, echo.Map{
		"data": user,
	})
}

func (ctrl *Controller) UpdateUserData(c echo.Context) (err error) {
	log.Print("enter ctrl.user.UpdateUserData")

	data := new(request.UpdateUserRequest)
	stringId := c.Param("id")
	err = c.Bind(&data)
	if err != nil {
		return c.JSON(400, &response.BasicUserResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	user, err := ctrl.service.Modify(stringId, data.DtoReq())
	if err != nil {
		return c.JSON(500, &response.BasicUserResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	return c.JSON(200, echo.Map{
		"data": user,
	})
}

func (ctrl *Controller) GetSingleData(c echo.Context) (err error) {
	log.Print("enter ctrl.user.GetSingleData")
	stringId := c.Param("id")
	user, err := ctrl.service.GetById(stringId)
	if err != nil {
		return c.JSON(500, &response.BasicUserResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	return c.JSON(200, echo.Map{
		"data": user,
	})
}

func (ctrl *Controller) DeleteData(c echo.Context) (err error) {
	log.Print("enter ctrl.user.DeleteData")
	stringId := c.Param("id")
	err = ctrl.service.Remove(stringId)
	if err != nil {
		return c.JSON(500, &response.BasicUserResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	return c.JSON(200, &response.BasicUserResponse{
		Status: "success",
	})
}
