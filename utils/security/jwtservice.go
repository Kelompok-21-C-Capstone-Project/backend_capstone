package security

import (
	"backend_capstone/models"
	"backend_capstone/services/user"
	"time"

	"github.com/golang-jwt/jwt"
)

type jwtService struct {
	key string
}

func NewJwtService(key string) user.JwtService {
	return &jwtService{
		key: key,
	}
}

func (service *jwtService) CreateJWT(data models.User) (tokenString string, err error) {
	timeNow := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, // header
		jwt.MapClaims{ // payload
			"id":         data.Id,
			"username":   data.Username,
			"name":       data.Name,
			"email":      data.Email,
			"phone":      data.Phone,
			"role":       []string{data.Role},
			"created_at": timeNow,
			"expired_at": timeNow.Add(1 * time.Hour),
		})
	tokenString, err = token.SignedString([]byte(service.key))
	if err != nil {
		return
	}
	return
}
