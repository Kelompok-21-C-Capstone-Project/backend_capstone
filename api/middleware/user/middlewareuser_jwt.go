package middleware

import (
	"backend_capstone/api/middleware/user/response"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type jwtUserMiddleware struct {
	key string
}

type JwtService interface {
	JwtUserMiddleware() echo.MiddlewareFunc
}

var (
	jwtSignedMethod = jwt.SigningMethodHS256
)

func NewJwtUserMiddleware(secretKey string) JwtService {
	return &jwtUserMiddleware{
		key: secretKey,
	}
}

func (s *jwtUserMiddleware) JwtUserMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			log.Print("enter middleware.JwtUserMiddleware")

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

			if reflect.ValueOf(claim["role"]).Index(0).Interface().(string) != "user" {
				return c.JSON(http.StatusForbidden, response.FailResponse{
					Status:  "fail",
					Message: "invalid token",
				})
			}

			c.Set("payload", fmt.Sprintf("%s", claim["id"]))

			return next(c)
		}
	}
}