package user

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

func (repo *PostgresRepository) FindById(id int) (transaction *models.User, err error) {
	return
}
func (repo *PostgresRepository) FindByQuery(key string, value interface{}) (products *[]models.User, err error) {
	return
}
func (repo *PostgresRepository) FindAll() (products *[]models.User, err error) {
	return
}
func (repo *PostgresRepository) Insert() (transaction *models.User, err error) {
	return
}
func (repo *PostgresRepository) Update() (transaction *models.User, err error) {
	return
}
func (repo *PostgresRepository) Delete() (transaction *models.User, err error) {
	return
}
