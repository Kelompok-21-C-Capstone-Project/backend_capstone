package user

import (
	"backend_capstone/api/user/request"
	"backend_capstone/api/user/response"
	"backend_capstone/services/user"
	"log"
	"net/http"

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

func (controller *Controller) GetAllData(c echo.Context) (err error) {
	log.Print("enter controller.user.GetAllData")
	// user, err := utils.ParsingJWT(c)
	// else if user.Role != "admin" {
	// 	return c.JSON(200, echo.Map{
	// 		"error": "restricted (*only for admin)",
	// 	})
	// }
	users, err := controller.service.GetAll()
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

func (controller *Controller) Create(c echo.Context) (err error) {
	log.Print("enter controller.user.Create")

	data := new(request.RegisterUserRequest)
	err = c.Bind(&data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &response.BasicUserResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	user, err := controller.service.Create(data.DtoReq())
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

func (controller *Controller) UpdateUserData(c echo.Context) (err error) {
	log.Print("enter controller.user.UpdateUserData")

	data := new(request.UpdateUserRequest)
	stringId := c.Param("id")
	err = c.Bind(&data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &response.BasicUserResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	user, err := controller.service.Modify(stringId, data.DtoReq())
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

func (controller *Controller) GetSingleData(c echo.Context) (err error) {
	log.Print("enter controller.user.GetSingleData")
	stringId := c.Param("id")
	user, err := controller.service.GetById(stringId)
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

func (controller *Controller) DeleteData(c echo.Context) (err error) {
	log.Print("enter controller.user.DeleteData")
	stringId := c.Param("id")
	payloadId := c.Get("payload").(string)
	err = controller.service.Remove(stringId, payloadId)
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

func (controller *Controller) AuthUser(c echo.Context) (err error) {
	log.Print("enter controller.user.AuthUser")
	data := new(request.LoginUserRequest)
	err = c.Bind(&data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &response.BasicUserResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	token, err := controller.service.UserLogin(data.DtoReq())
	if err != nil {
		return c.JSON(http.StatusUnauthorized, &response.BasicUserResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusAccepted, &response.BasicUserResponse{
		Status:  "success",
		Message: token,
	})
}
