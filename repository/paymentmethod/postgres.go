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
	if err = repo.db.Preload("PaymentVendors").First(&paymentMethod, id).Error; err != nil {
		return
	}
	return
}
func (repo *PostgresRepository) FindByQuery(key string, value interface{}) (paymentMethods *[]models.PaymentMethod, err error) {
	return
}
func (repo *PostgresRepository) FindAll() (paymentMethods *[]models.PaymentMethod, err error) {
	if err = repo.db.Preload("PaymentVendors").Find(&paymentMethods).Error; err != nil {
		return
	}
	return
}
func (repo *PostgresRepository) Insert(data *models.PaymentMethod) (paymentMethod *models.PaymentMethod, err error) {
	if err = repo.db.Create(data).Error; err != nil {
		return
	}
	return data, err
}
func (repo *PostgresRepository) Update(id string, data *models.PaymentMethod) (paymentMethod *models.PaymentMethod, err error) {
	if err = repo.db.First(&paymentMethod, id).Model(paymentMethod).Updates(data).Preload("PaymentVendors").First(&paymentMethod, id).Error; err != nil {
		return
	}
	return
}
func (repo *PostgresRepository) Delete(id string) (err error) {
	if err = repo.db.Delete(&models.PaymentMethod{}, id).Error; err != nil {
		return
	}
	return
}
