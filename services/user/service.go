package user

import (
	"backend_capstone/models"
	"backend_capstone/services/user/dto"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Repository interface {
	FindById(id string) (user *models.User, err error)
	FindByQuery(key string, value interface{}) (users *[]models.User, err error)
	FindAll() (users *[]models.User, err error)
	Insert(data *models.User) (user *models.User, err error)
	Update(id string, data *models.User) (user *models.User, err error)
	Delete(id string) (err error)
}

type Service interface {
	GetById(id string) (user models.User, err error)
	GetAll() (users []models.User, err error)
	Create(registeruserDTO dto.RegisterUserDTO) (user models.User, err error)
	Modify(id string, updateuserDTO dto.UpdateUserDTO) (user models.User, err error)
	Remove(id string) (err error)
}

type service struct {
	repository Repository
	validate   *validator.Validate
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
		validate:   validator.New(),
	}
}

// Untuk mengambil data user berdasarkan id
func (s *service) GetById(id string) (user models.User, err error) {
	_, err = uuid.Parse(id)
	if err != nil {
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
func (s *service) GetAll() (users []models.User, err error) {
	datas, err := s.repository.FindAll()
	if err != nil {
		return
	}
	users = *datas
	return
}

// Registrasi user
func (s *service) Create(registeruserDTO dto.RegisterUserDTO) (user models.User, err error) {
	err = s.validate.Struct(registeruserDTO)
	if err != nil {
		return
	}
	id := uuid.New().String()
	data, err := s.repository.Insert(registeruserDTO.GenerateModel(id))
	if err != nil {
		return
	}
	user = *data
	return
}
func (s *service) Modify(id string, updateuserDTO dto.UpdateUserDTO) (user models.User, err error) {
	_, err = uuid.Parse(id)
	if err != nil {
		return
	}
	_, err = s.repository.FindById(id)
	if err != nil {
		return
	}
	s.repository.Update(id, updateuserDTO.GenerateModel())
	return
}
func (s *service) Remove(id string) (err error) {
	_, err = uuid.Parse(id)
	if err != nil {
		return
	}
	err = s.repository.Delete(id)
	if err != nil {
		return
	}
	return
}
