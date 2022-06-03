package products

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

func (repo *PostgresRepository) FindById(id int) (transaction *models.Product, err error) {
	return
}
func (repo *PostgresRepository) FindByQuery(key string, value interface{}) (products *[]models.Product, err error) {
	return
}
func (repo *PostgresRepository) FindAll() (products *[]models.Product, err error) {
	return
}
func (repo *PostgresRepository) Insert() (transaction *models.Product, err error) {
	return
}
func (repo *PostgresRepository) Update() (transaction *models.Product, err error) {
	return
}
func (repo *PostgresRepository) Delete() (transaction *models.Product, err error) {
	return
}
