package productcategory

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

func (repo *PostgresRepository) FindById(id string) (productCategory *models.ProductCategory, err error) {
	return
}
func (repo *PostgresRepository) FindByQuery(key string, value interface{}) (productCategories *[]models.ProductCategory, err error) {
	return
}
func (repo *PostgresRepository) FindAll() (productCategories *[]models.ProductCategory, err error) {
	return
}
func (repo *PostgresRepository) Insert() (productCategory *models.ProductCategory, err error) {
	return
}
func (repo *PostgresRepository) Update() (productCategory *models.ProductCategory, err error) {
	return
}
func (repo *PostgresRepository) Delete() (productCategory *models.ProductCategory, err error) {
	return
}
