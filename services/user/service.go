package user

import (
	"backend_capstone/models"
	"backend_capstone/services/user/dto"
	"errors"
	"math"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Mailjet interface {
	ValidateEmail() (err error)
	RecoverPassword() (err error)
}

type JwtService interface {
	CreateJWT(data models.User) (token string, err error)
}

type PasswordHash interface {
	Hash(password string) (hash string, err error)
	CheckPassword(password string, hash string) (err error)
}

type Repository interface {
	FindById(id string) (user *models.UserResponse, err error)
	FindByIdentifier(identifier string) (user *models.User, err error)
	FindAll() (users *[]models.UserResponse, err error)
	Insert(data *models.User) (user *models.UserResponse, err error)
	Update(id string, data *models.User) (user *models.UserResponse, err error)
	Delete(id string) (err error)
	DashboardDetailUser(params ...string) (data dto.UserDashboadDTO, err error)
}

type Service interface {
	GetById(id string, payloadId string) (user models.UserResponse, err error)
	GetAll() (users []models.UserResponse, err error)
	Create(registeruserDTO dto.RegisterUserDTO) (user models.UserResponse, err error)
	CreateAdmin(registeradminDTO dto.RegisterAdminDTO) (user models.UserResponse, err error)
	Modify(id string, payloadId string, updateuserDTO dto.UpdateUserDTO) (user models.UserResponse, err error)
	Remove(id string, payloadId string) (err error)
	UserLogin(loginuserDTO dto.LoginUserDTO) (token string, err error)
	AdminDetailUser(params ...string) (data dto.UserDashboadDTO, err error)
}

type service struct {
	repository Repository
	hasher     PasswordHash
	validate   *validator.Validate
	jwtService JwtService
	adminKey   string
}

func NewService(repository Repository, hasher PasswordHash, jwtService JwtService, adminKey string) Service {
	return &service{
		repository: repository,
		hasher:     hasher,
		jwtService: jwtService,
		adminKey:   adminKey,
		validate:   validator.New(),
	}
}

func (s *service) AdminDetailUser(params ...string) (data dto.UserDashboadDTO, err error) {
	data, err = s.repository.DashboardDetailUser(params...)
	if params[1] == "" {
		params[1] = "1"
	}
	if params[2] == "" {
		params[2] = "10"
	}
	den, err := strconv.Atoi(params[2])
	if err != nil {
		return
	}
	if den < -1 {
		data.PageLength = int64(math.Ceil(float64(data.PageLength) / float64(10)))
	} else if den == -1 || den == 0 {
		data.PageLength = 1
	} else {
		data.PageLength = int64(math.Ceil(float64(data.PageLength) / float64(den)))
	}
	if err != nil {
		return
	}
	return
}

// Untuk mengambil data user berdasarkan id
func (s *service) GetById(id string, payloadId string) (user models.UserResponse, err error) {
	_, err = uuid.Parse(id)
	if err != nil {
		return
	}
	if id != payloadId {
		err = errors.New("unauthorized")
		return
	}
	data, err := s.repository.FindById(id)
	if err != nil {
		return
	}
	user = *data
	return
}

// Untuk mengambil semua data user
func (s *service) GetAll() (users []models.UserResponse, err error) {
	datas, err := s.repository.FindAll()
	if err != nil {
		return
	}
	users = *datas
	return
}

// Registrasi admin
func (s *service) CreateAdmin(registeradminDTO dto.RegisterAdminDTO) (user models.UserResponse, err error) {
	if registeradminDTO.Key != s.adminKey {
		err = errors.New("unauthorized")
		return
	}
	err = s.validate.Struct(registeradminDTO)
	if err != nil {
		return
	}
	id := uuid.New().String()
	registeradminDTO.Password, err = s.hasher.Hash(registeradminDTO.Password)
	if err != nil {
		return
	}
	data, err := s.repository.Insert(registeradminDTO.GenerateModel(id))
	if err != nil {
		return
	}
	user = *data
	return
}

// Registrasi user
func (s *service) Create(registeruserDTO dto.RegisterUserDTO) (user models.UserResponse, err error) {
	err = s.validate.Struct(registeruserDTO)
	if err != nil {
		return
	}
	id := uuid.New().String()
	registeruserDTO.Password, err = s.hasher.Hash(registeruserDTO.Password)
	if err != nil {
		return
	}
	data, err := s.repository.Insert(registeruserDTO.GenerateModel(id))
	if err != nil {
		return
	}
	user = *data
	return
}
func (s *service) Modify(id string, payloadId string, updateuserDTO dto.UpdateUserDTO) (user models.UserResponse, err error) {
	_, err = uuid.Parse(id)
	if err != nil {
		return
	}
	_, err = s.repository.FindById(id)
	if err != nil {
		return
	}
	if id != payloadId {
		err = errors.New("unauthorized")
		return
	}
	updateuserDTO.Password, err = s.hasher.Hash(updateuserDTO.Password)
	if err != nil {
		return
	}
	data, err := s.repository.Update(id, updateuserDTO.GenerateModel())
	user = *data
	return
}
func (s *service) Remove(id string, payloadId string) (err error) {
	_, err = uuid.Parse(id)
	if err != nil {
		return
	}
	err = s.repository.Delete(id)
	if err != nil {
		return
	}
	if id != payloadId {
		return errors.New("unauthorized")
	}
	return
}
func (s *service) UserLogin(loginuserDTO dto.LoginUserDTO) (token string, err error) {
	err = s.validate.Struct(loginuserDTO)
	if err != nil {
		return
	}
	user, err := s.repository.FindByIdentifier(loginuserDTO.Identifier)
	if err != nil {
		return
	}
	err = s.hasher.CheckPassword(loginuserDTO.Password, user.Password)
	if err != nil {
		return
	}
	// ngegenerate token jwt
	token, err = s.jwtService.CreateJWT(*user)
	if err != nil {
		return
	}
	return
}
