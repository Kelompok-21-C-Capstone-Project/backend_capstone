package payments

import (
	"backend_capstone/models"
	"backend_capstone/utils/midtransdriver"
	"log"

	"github.com/go-playground/validator/v10"
)

type Repository interface {
	FindById(id int) (payment *models.Payment, err error)
	FindByQuery(key string, value interface{}) (payments *[]models.Payment, err error)
	FindAll() (payments *[]models.Payment, err error)
	Insert() (payment *models.Payment, err error)
	Update() (payment *models.Payment, err error)
	Delete() (payment *models.Payment, err error)
}

type Service interface {
	GetById(id int) (payment models.Payment, err error)
	GetAll() (payments []models.Payment, err error)
	Create() (payment models.Payment, err error)
	Modify() (payment models.Payment, err error)
	Remove() (err error)
}

type service struct {
	repository Repository
	validate   *validator.Validate
	midtrans   *midtransdriver.MidtransDriver
}

func NewService(repository Repository, midtransApi *midtransdriver.MidtransDriver) Service {
	log.Print("Enter NewService")
	return &service{
		repository: repository,
		validate:   validator.New(),
		midtrans:   midtransApi,
	}
}

func (s *service) GetById(id int) (payment models.Payment, err error) {
	return
}
func (s *service) GetAll() (payments []models.Payment, err error) {
	return
}
func (s *service) Create() (payment models.Payment, err error) {
	log.Print("Enter PaymentService.Create")
	s.midtrans.GetPaymentMethod()
	return
}
func (s *service) Modify() (payment models.Payment, err error) {
	log.Print("Enter PaymentService.Modify")
	s.midtrans.PutApprovePaymentMethod()
	return
}
func (s *service) Remove() (err error) {
	return
}
