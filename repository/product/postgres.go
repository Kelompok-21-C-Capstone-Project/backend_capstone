package products

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

func (repo *PostgresRepository) FindById(id uuid.UUID) (product *models.Product, err error) {
	if err = repo.db.First(&product, id).Error; err != nil {
		return
	}
	return
}
func (repo *PostgresRepository) FindByQuery(key string, value interface{}) (products *[]models.Product, err error) {
	return
}
func (repo *PostgresRepository) FindAll() (products *[]models.Product, err error) {
	if err = repo.db.Find(&products).Error; err != nil {
		return
	}
	return
}
func (repo *PostgresRepository) Insert(data *models.Product) (product *models.Product, err error) {
	if err = repo.db.Create(&product).Error; err != nil {
		return
	}
	return
}
func (repo *PostgresRepository) Update(id uuid.UUID, data *models.Product) (product *models.Product, err error) {
	if err = repo.db.First(&product, id).Model(product).Updates(data).Error; err != nil {
		return
	}
	return
}
func (repo *PostgresRepository) Delete(id uuid.UUID) (product *models.Product, err error) {
	return
}
func (repo *PostgresRepository) ValidateProductBrandCategories(brandId uuid.UUID, categoryId uuid.UUID) (productBrandCategoriesId uuid.UUID, err error) {
	check := models.ProductBrandCategory{}
	rowCount := repo.db.Where("product_brand_id = ? AND product_category_id = ?", brandId, categoryId).Find(&check).RowsAffected
	if rowCount > 0 {
		return check.Id, err
	}
	id := uuid.New()
	if err = repo.db.Create(&models.ProductBrandCategory{
		Id:                id,
		ProductBrandId:    brandId,
		ProductCategoryId: categoryId,
		Slug:              "",
		IsAvailable:       false,
	}).Error; err != nil {
		return
	}
	return id, err
}
