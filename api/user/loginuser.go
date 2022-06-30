package user

import (
	"backend_capstone/models"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"backend_capstone/helpers"

	"github.com/labstack/echo"
)

type Admin struct {
	Email    string `json:"username" bson:"username" validate:"required,email"`
	Password string `json:"password,omitempty" bson:"password" validate:"required,min=8,max=300"`
	IsAdmin  bool   `json:"isadmin,omitempty" bson:"isadmin"`
}

func CheckLogin(c echo.Context) error {
	Username := c.FormValue("username")
	Password := c.FormValue("password")

	res, err := models.User(Username, Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"messages": err.Error(),
		})
	}

	if !res {
		return echo.ErrUnauthorized
	}

	// token jwt
	func (u Admin) createToken() (string, error) {
		if err := cleanenv.ReadEnv(&prop); err != nil {
			log.Errorf("Configuration cannot be read : %v", err)
		}

		claims := jwt.MapClaims{}
		claims["authorized"] = u.IsAdmin
		claims["user_id"] = u.Email
		claims["exp"] = time.Now().Add(time.Minute * 15).Unix()
		at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		token, err := at.SignedString([]byte(prop.JwtTokenSecret))
		if err != nil {
			log.Errorf("Unable to generate the token :%v", err)
			return "", err
		}
		return token, nil
	}

	// t, err := token.SignedString([]byte("uuid")) //belum diperlukan
	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, map[string]string{
	// 		"messages": err.Error(),
	// 	})
	// }

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}


func GenerateHashPassword(c echo.Context) error {
	password := c.Param("password")

	hash, _ := helpers.HashPassword(password)

	return c.JSON(http.StatusOK, hash)
}



