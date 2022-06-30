package productcategory

import (
	"backend_capstone/models"

	"github.com/google/uuid"
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

func (repo *PostgresRepository) FindById(id uuid.UUID) (productCategory *models.ProductCategory, err error) {
	if err = repo.db.First(&productCategory, id).Error; err != nil {
		return
	}
	return
}
func (repo *PostgresRepository) FindByQuery(key string, value interface{}) (productCategories *[]models.ProductCategory, err error) {
	return
}
func (repo *PostgresRepository) FindAll() (productCategories *[]models.ProductCategory, err error) {
	if err = repo.db.Find(&productCategories).Error; err != nil {
		return
	}
	return
}
func (repo *PostgresRepository) Insert(data *models.ProductCategory) (productCategory *models.ProductCategory, err error) {
	if err = repo.db.Create(data).Error; err != nil {
		return
	}
	return data, err
}
func (repo *PostgresRepository) Update(id uuid.UUID, data *models.ProductCategory) (productCategory *models.ProductCategory, err error) {
	if err = repo.db.First(&productCategory, id).Model(productCategory).Updates(data).Error; err != nil {
		return
	}
	return
}
func (repo *PostgresRepository) Delete(id uuid.UUID) (err error) {
	if err = repo.db.Delete(&models.ProductCategory{}, id).Error; err != nil {
		return
	}
	return
}
