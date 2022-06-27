package paymentmethod

import (
	"backend_capstone/models"

	"gorm.io/gorm"
)

type PostgresRepository struct {
	db *gorm.DB
}

func NewPostgresRepository(db *gorm.DB) *PostgresRepository {
	return &PostgresRepository{
		db: db,
	}
}

func (repo *PostgresRepository) FindById(id string) (paymentMethod *models.PaymentMethod, err error) {
	return
}
func (repo *PostgresRepository) FindByQuery(key string, value interface{}) (paymentMethods *[]models.PaymentMethod, err error) {
	return
}
func (repo *PostgresRepository) FindAll() (paymentMethods *[]models.PaymentMethod, err error) {
	return
}
func (repo *PostgresRepository) Insert() (paymentMethod *models.PaymentMethod, err error) {
	return
}
func (repo *PostgresRepository) Update() (paymentMethod *models.PaymentMethod, err error) {
	return
}
func (repo *PostgresRepository) Delete() (paymentMethod *models.PaymentMethod, err error) {
	return
}
