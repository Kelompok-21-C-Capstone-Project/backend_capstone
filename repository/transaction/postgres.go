package transaction

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

func (repo *PostgresRepository) FindById(id string) (transaction *models.Transaction, err error) {
	return
}
func (repo *PostgresRepository) FindByQuery(key string, value interface{}) (transactions *[]models.Transaction, err error) {
	return
}
func (repo *PostgresRepository) FindAll() (transactions *[]models.Transaction, err error) {
	return
}
func (repo *PostgresRepository) Insert(data *models.Transaction) (transaction *models.Transaction, err error) {
	if err = repo.db.Create(&data).Error; err != nil {
		return
	}
	return data, err
}
func (repo *PostgresRepository) InsertPayment(data *models.Payment) (transaction *models.Payment, err error) {
	if err = repo.db.Create(&data).Error; err != nil {
		return
	}
	return data, err
}
func (repo *PostgresRepository) Update() (transaction *models.Transaction, err error) {
	return
}
func (repo *PostgresRepository) Delete() (transaction *models.Transaction, err error) {
	return
}
