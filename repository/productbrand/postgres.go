package productbrand

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

func (repo *PostgresRepository) FindById(id string) (productBrand *models.ProductBrand, err error) {
	return
}
func (repo *PostgresRepository) FindByQuery(key string, value interface{}) (productBrands *[]models.ProductBrand, err error) {
	return
}
func (repo *PostgresRepository) FindAll() (productBrands *[]models.ProductBrand, err error) {
	return
}
func (repo *PostgresRepository) Insert() (productBrand *models.ProductBrand, err error) {
	return
}
func (repo *PostgresRepository) Update() (productBrand *models.ProductBrand, err error) {
	return
}
func (repo *PostgresRepository) Delete() (productBrand *models.ProductBrand, err error) {
	return
}
func (repo *PostgresRepository) GetAllCategory() (productBrands *[]models.ProductCategory, err error) {
	return
}
func (repo *PostgresRepository) InsertBrandCategory() (productBrand *models.ProductBrandCategory, err error) {
	return
}
func (repo *PostgresRepository) DeleteBrandCategory() (err error) {
	return
}
