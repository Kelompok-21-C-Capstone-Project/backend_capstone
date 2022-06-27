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
func (repo *PostgresRepository) Insert() (transaction *models.Transaction, err error) {
	return
}
func (repo *PostgresRepository) Update() (transaction *models.Transaction, err error) {
	return
}
func (repo *PostgresRepository) Delete() (transaction *models.Transaction, err error) {
	return
}
