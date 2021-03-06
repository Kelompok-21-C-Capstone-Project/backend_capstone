package admin

import (
	"backend_capstone/api/middleware/admin/response"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type jwtAdminMiddleware struct {
	key string
}

type JwtService interface {
	JwtAdminMiddleware() echo.MiddlewareFunc
}

var (
	jwtSignedMethod = jwt.SigningMethodHS256
)

func NewJwtAdminMiddleware(secretKey string) JwtService {
	return &jwtAdminMiddleware{
		key: secretKey,
	}
}

func (s *jwtAdminMiddleware) JwtAdminMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			log.Print("enter middleware.JwtAdminMiddleware")

			signature := strings.Split(c.Request().Header.Get("Authorization"), " ")
			if len(signature) < 2 {
				return c.JSON(http.StatusForbidden, response.FailResponse{
					Status:  "fail",
					Message: "invalid token",
				})
			}

			if signature[0] != "Bearer" {
				return c.JSON(http.StatusForbidden, response.FailResponse{
					Status:  "fail",
					Message: "invalid token",
				})
			}

			claim := jwt.MapClaims{}

			token, _ := jwt.ParseWithClaims(signature[1], claim, func(t *jwt.Token) (interface{}, error) {
				return []byte(s.key), nil
			})

			expiredAt, _ := time.Parse(time.RFC3339, claim["expired_at"].(string))
			if expiredAt.Before(time.Now()) {
				return c.JSON(http.StatusRequestTimeout, response.FailResponse{
					Status:  "fail",
					Message: "Token expired",
				})
			}

			method, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok || method != jwtSignedMethod {
				return c.JSON(http.StatusForbidden, response.FailResponse{
					Status:  "fail",
					Message: "invalid token",
				})
			}
			if reflect.ValueOf(claim["role"]).Index(0).Interface().(string) != "admin" {
				return c.JSON(http.StatusForbidden, response.FailResponse{
					Status:  "fail",
					Message: "invalid token",
				})
			}

			c.Set("payload", fmt.Sprintf("%s", claim["id"]))
			c.Set("username", fmt.Sprintf("%s", claim["username"]))
			c.Set("name", fmt.Sprintf("%s", claim["name"]))
			c.Set("email", fmt.Sprintf("%s", claim["email"]))
			c.Set("phone", fmt.Sprintf("%s", claim["phone"]))
			c.Set("role", fmt.Sprintf("%s", reflect.ValueOf(claim["role"]).Index(0).Interface().(string)))

			return next(c)
		}
	}
}
