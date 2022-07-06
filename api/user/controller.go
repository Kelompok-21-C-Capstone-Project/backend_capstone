package user

import (
	"backend_capstone/api/user/request"
	"backend_capstone/api/user/response"
	"backend_capstone/services/user"
	"log"
	"net/http"
	"strings"
	"time"

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

// Create godoc
// @Summary User register
// @Description  Create new user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param Payload body request.RegisterUserRequest true "Payload format" SchemaExample(request.RegisterUserRequest)
// @Success      201  {object}  models.UserResponse
// @Failure      400  {object}  response.BasicUserResponse
// @Failure      403  {object}  response.BasicUserResponse
// @Failure      500  {object}  response.BasicUserResponse
// @Router       /v1/user_register [post]
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

// CreateAdmin godoc
// @Summary Admin register
// @Description  Create new admin
// @Tags         admins
// @Accept       json
// @Produce      json
// @Param Payload body request.RegisterAdminRequest true "Payload format" SchemaExample(request.RegisterAdminRequest)
// @Success      201  {object}  models.UserResponse
// @Failure      400  {object}  response.BasicUserResponse
// @Failure      403  {object}  response.BasicUserResponse
// @Failure      500  {object}  response.BasicUserResponse
// @Router       /v1/admin_register [post]
func (controller *Controller) CreateAdmin(c echo.Context) (err error) {
	log.Print("enter controller.user.Create")

	data := new(request.RegisterAdminRequest)
	err = c.Bind(&data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &response.BasicUserResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	user, err := controller.service.CreateAdmin(data.DtoReq())
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

// UpdateUserData godoc
// @Summary Update user
// @Description  Update user data
// @Tags         users
// @Accept       json
// @Produce      json
// @Param id   path  string  true  "User ID" minLength:"32"
// @Param Payload body request.UpdateUserRequest true "Payload format" SchemaExample(request.UpdateUserRequest)
// @Success      200  {object}  models.UserResponse
// @Failure      400  {object}  response.BasicUserResponse
// @Failure      403  {object}  response.BasicUserResponse
// @Failure      500  {object}  response.BasicUserResponse
// @Security ApiKeyAuth
// @Router       /v1/users/{id} [put]
func (controller *Controller) UpdateUserData(c echo.Context) (err error) {
	log.Print("enter controller.user.UpdateUserData")

	data := new(request.UpdateUserRequest)
	stringId := c.Param("id")
	payloadId := c.Get("payload").(string)
	err = c.Bind(&data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &response.BasicUserResponse{
			Status:  "fail",
			Message: err.Error(),
		})
	}
	user, err := controller.service.Modify(stringId, payloadId, data.DtoReq())
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

// GetSingleData godoc
// @Summary Get user data by id
// @Description  Get detailed user data by id from database
// @Tags         users
// @Produce      json
// @Param id   path  string  true  "User ID" minLength:"32"
// @Success      200  {object}  models.UserResponse
// @Failure      400  {object}  response.BasicUserResponse
// @Failure      403  {object}  response.BasicUserResponse
// @Failure      500  {object}  response.BasicUserResponse
// @Security ApiKeyAuth
// @Router       /v1/users/{id} [get]
func (controller *Controller) GetSingleData(c echo.Context) (err error) {
	log.Print("enter controller.user.GetSingleData")
	stringId := c.Param("id")
	payloadId := c.Get("payload").(string)
	user, err := controller.service.GetById(stringId, payloadId)
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

// DeleteData godoc
// @Summary Delete user data by id
// @Description  Delete user data from database
// @Tags         users
// @Produce      json
// @Param id   path  string  true  "User ID" minLength:"32"
// @Success      200  {object}  response.BasicUserSuccessResponse
// @Failure      400  {object}  response.BasicUserResponse
// @Failure      403  {object}  response.BasicUserResponse
// @Failure      500  {object}  response.BasicUserResponse
// @Security ApiKeyAuth
// @Router       /v1/users/{id} [delete]
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
	return c.JSON(200, &response.BasicUserSuccessResponse{
		Status: "success",
	})
}

// DeleteData godoc
// @Summary Delete user data by id
// @Description  Delete user data from database
// @Tags         users
// @Accept       json
// @Produce      json
// @Param Payload body request.LoginUserRequest true "Payload format" SchemaExample(request.LoginUserRequest)
// @Success      200  {object}  response.SuccessLoginResponse
// @Failure      400  {object}  response.BasicUserResponse
// @Failure      403  {object}  response.BasicUserResponse
// @Failure      500  {object}  response.BasicUserResponse
// @Security ApiKeyAuth
// @Router       /v1/auth [post]
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
	return c.JSON(http.StatusAccepted, &response.SuccessLoginResponse{
		Status: "success",
		Token:  token,
	})
}

// GetSingleData godoc
// @Summary Get token payload
// @Description  Get detailed token data
// @Tags         tokens
// @Produce      json
// @Success      200  {object}  response.JwtPayload
// @Failure      400  {object}  response.BasicUserResponse
// @Failure      403  {object}  response.BasicUserResponse
// @Failure      500  {object}  response.BasicUserResponse
// @Security ApiKeyAuth
// @Router       /v1/tokens [get]
func (controller *Controller) ParseToken(c echo.Context) (err error) {
	log.Print("enter controller.user.ParseToken")
	signature := strings.Split(c.Request().Header.Get("Authorization"), " ")
	if len(signature) < 2 {
		log.Print("fail")
		return c.JSON(http.StatusForbidden, response.BasicUserResponse{
			Status:  "fail",
			Message: "invalid token",
		})
	}

	if signature[0] != "Bearer" {
		log.Print("fail 2")
		return c.JSON(http.StatusForbidden, response.BasicUserResponse{
			Status:  "fail",
			Message: "invalid token",
		})
	}
	iatTime, _ := time.Parse(time.RFC3339, c.Get("created_at").(string))
	expTime, _ := time.Parse(time.RFC3339, c.Get("expired_at").(string))
	return c.JSON(200, &response.JwtPayload{
		Id:       c.Get("payload").(string),
		Username: c.Get("username").(string),
		Name:     c.Get("name").(string),
		Role:     []string{c.Get("role").(string)},
		Email:    c.Get("email").(string),
		Phone:    c.Get("phone").(string),
		Iat:      iatTime.UnixMilli(),
		Exp:      expTime.UnixMilli(),
	})
}
