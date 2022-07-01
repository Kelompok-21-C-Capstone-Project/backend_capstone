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
	if err = repo.db.First(&paymentVendor, id).Error; err != nil {
		return
	}
	return
}
func (repo *PostgresRepository) FindByQuery(key string, value interface{}) (paymentVendors *[]models.PaymentVendor, err error) {
	return
}
func (repo *PostgresRepository) FindAll() (paymentVendors *[]models.PaymentVendor, err error) {
	if err = repo.db.Find(&paymentVendors).Error; err != nil {
		return
	}
	return
}
func (repo *PostgresRepository) ValidatePaymentMethod(MethodId string) (err error) {
	if err = repo.db.Table("payment_methods").First(&models.PaymentMethod{}, MethodId).Error; err != nil {
		return
	}
	return
}
func (repo *PostgresRepository) Insert(data *models.PaymentVendor) (paymentVendor *models.PaymentVendor, err error) {
	if err = repo.db.Create(&data).Error; err != nil {
		return
	}
	return data, err
}
func (repo *PostgresRepository) Update(id string, data *models.PaymentVendor) (paymentVendor *models.PaymentVendor, err error) {
	if err = repo.db.First(&paymentVendor, id).Model(paymentVendor).Updates(data).First(&paymentVendor, id).Error; err != nil {
		return
	}
	return
}
func (repo *PostgresRepository) Delete(id string) (paymentVendor *models.PaymentVendor, err error) {
	if err = repo.db.Delete(&paymentVendor, id).Error; err != nil {
		return
	}
	return
}
