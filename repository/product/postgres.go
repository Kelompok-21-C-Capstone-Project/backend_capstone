package products

import (
	"backend_capstone/models"
	"backend_capstone/services/product/dto"

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

func (repo *PostgresRepository) FindById(id string) (product *models.Product, err error) {
	if err = repo.db.Preload("ProductBrandCategory").First(&product, &id).Error; err != nil {
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
func (repo *PostgresRepository) ClientFindAllBySlug(slug string) (products *dto.ProductCategory, err error) {
	err = repo.db.Table("product_categories").Select("product_categories.id, product_categories.name, product_categories.slug").Find(&products, "slug = ?", &slug).Error
	if err != nil {
		return
	}
	if err = repo.db.Table("product_brand_categories").Select("products.id, product_brands.name as group, products.name as label, products.description as description, products.stock as stock, products.price as price, products.is_discount").Joins("left join product_categories on product_brand_categories.product_category_id = product_categories.id").Joins("left join products on product_brand_categories.id = products.product_brand_category_id").Joins("left join product_brands on product_brand_categories.product_brand_id = product_brands.id").Where("product_brand_categories.product_category_id = ?", products.Id).Find(&products.Products).Error; err != nil {
		return
	}
	return
}
func (repo *PostgresRepository) ClientFindAll() (products *[]dto.ProductCategory, err error) {
	err = repo.db.Table("product_categories").Select("product_categories.id, product_categories.name, product_categories.slug").Find(&products).Error
	if err != nil {
		return
	}
	for i, el := range *products {
		if err = repo.db.Table("product_brand_categories").Select("products.id, product_brands.name as group, products.name as label, products.description as description, products.stock as stock, products.price as price, products.is_discount").Joins("left join product_categories on product_brand_categories.product_category_id = product_categories.id").Joins("left join products on product_brand_categories.id = products.product_brand_category_id").Joins("left join product_brands on product_brand_categories.product_brand_id = product_brands.id").Where("product_brand_categories.product_category_id = ?", el.Id).Find(&(*products)[i].Products).Error; err != nil {
			return
		}
	}
	return
}
func (repo *PostgresRepository) Insert(data *models.Product) (product *models.Product, err error) {
	if err = repo.db.Create(&data).Error; err != nil {
		return
	}
	return data, err
}
func (repo *PostgresRepository) Update(id string, data *models.Product) (product *models.Product, err error) {
	if err = repo.db.First(&product, &id).Model(product).Updates(data).Preload("ProductBrandCategory").First(&product, &id).Error; err != nil {
		return
	}
	return
}
func (repo *PostgresRepository) Delete(id string) (err error) {
	if err = repo.db.Delete(&models.Product{}, &id).Error; err != nil {
		return
	}
	return
}
func (repo *PostgresRepository) ValidateProductBrandCategories(brandId string, categoryId string) (productBrandCategoriesId string, err error) {
	check := models.ProductBrandCategory{}
	rowCount := repo.db.Where("product_brand_id = ? AND product_category_id = ?", brandId, categoryId).Find(&check).RowsAffected
	if rowCount > 0 {
		return check.Id, err
	}
	id := uuid.New()
	if err = repo.db.Create(&models.ProductBrandCategory{
		Id:                id.String(),
		ProductBrandId:    brandId,
		ProductCategoryId: categoryId,
		Slug:              "",
		IsAvailable:       false,
	}).Error; err != nil {
		return
	}
	return id.String(), err
}
