package paymentvendor

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

func (repo *PostgresRepository) FindById(id string) (paymentVendor *models.PaymentVendor, err error) {
	return
}
func (repo *PostgresRepository) FindByQuery(key string, value interface{}) (paymentVendors *[]models.PaymentVendor, err error) {
	return
}
func (repo *PostgresRepository) FindAll() (paymentVendors *[]models.PaymentVendor, err error) {
	return
}
func (repo *PostgresRepository) Insert() (paymentVendor *models.PaymentVendor, err error) {
	return
}
func (repo *PostgresRepository) Update() (paymentVendor *models.PaymentVendor, err error) {
	return
}
func (repo *PostgresRepository) Delete() (paymentVendor *models.PaymentVendor, err error) {
	return
}
