package productbrand

import (
	"backend_capstone/models"
	"errors"
	"log"

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

func (repo *PostgresRepository) FindById(id string) (productBrand *models.ProductBrand, err error) {
	uid, err := uuid.Parse(id)
	if err = repo.db.Preload("ProductCategories").First(&productBrand, uid).Error; err != nil {
		return
	}
	return
}
func (repo *PostgresRepository) FindByQuery(key string, value interface{}) (productBrands *[]models.ProductBrand, err error) {
	return
}
func (repo *PostgresRepository) FindAll() (productBrands *[]models.ProductBrand, err error) {
	if err = repo.db.Preload("ProductCategories").Find(&productBrands).Error; err != nil {
		return
	}
	return
}
func (repo *PostgresRepository) FindCategoryById(id string) (productCategory *models.ProductCategory, err error) {
	uid, err := uuid.Parse(id)
	if err = repo.db.First(&productCategory, uid).Error; err != nil {
		return
	}
	return
}
func (repo *PostgresRepository) Insert(data *models.ProductBrand) (productBrand *models.ProductBrand, err error) {
	if err = repo.db.Create(data).Preload("ProductCategories").First(&productBrand, data.Id).Error; err != nil {
		return
	}
	return
}
func (repo *PostgresRepository) Update(id string, data *models.ProductBrand) (productBrand *models.ProductBrand, err error) {
	uid, err := uuid.Parse(id)
	if err = repo.db.First(&productBrand, uid).Model(productBrand).Updates(data).Error; err != nil {
		return
	}
	log.Print(productBrand)
	return
}
func (repo *PostgresRepository) Delete(id string) (err error) {
	uid, err := uuid.Parse(id)
	if err = repo.db.Delete(&models.ProductBrand{}, uid).Error; err != nil {
		return
	}
	return
}
func (repo *PostgresRepository) CheckBrandCategory(brandId uuid.UUID, categoryId uuid.UUID) (rowCount int64, err error) {
	rowCount = repo.db.Where("product_brand_id = ? AND product_category_id = ?", brandId, categoryId).Find(&models.ProductBrandCategory{}).RowsAffected
	if rowCount > 0 {
		return rowCount, errors.New("")
	}
	return
}
func (repo *PostgresRepository) InsertBrandCategory(brandId uuid.UUID, categoryId uuid.UUID, slug string) (productBrand *models.ProductBrandCategory, err error) {
	productBrand = &models.ProductBrandCategory{
		Id:                uuid.New(),
		ProductBrandId:    brandId,
		ProductCategoryId: categoryId,
		Slug:              slug,
		IsAvailable:       false,
	}
	if err = repo.db.Create(productBrand).Error; err != nil {
		return
	}
	return
}
func (repo *PostgresRepository) DeleteBrandCategory(brandId uuid.UUID, categoryId uuid.UUID) (err error) {
	if err = repo.db.Where("product_brand_id = ? AND product_category_id = ?", brandId, categoryId).Delete(&models.ProductBrandCategory{}).Error; err != nil {
		return
	}
	return
}
