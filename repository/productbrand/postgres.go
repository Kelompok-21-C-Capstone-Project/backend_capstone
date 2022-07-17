package productbrand

import (
	"backend_capstone/models"
	"backend_capstone/services/productbrand/dto"
	"errors"
	"log"
	"strconv"

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
func Paginate(page string, pageSize string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(page)
		if page == 0 {
			page = 1
		}

		pageSize, _ := strconv.Atoi(pageSize)
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func (repo *PostgresRepository) FindById(id string) (productBrand *models.ProductBrandResponse, err error) {
	if err = repo.db.First(&models.ProductBrand{}, &id).Scan(&productBrand).Error; err != nil {
		return
	}
	if err = repo.db.Debug().Table("product_categories").Joins("left join product_brand_categories on product_brand_categories.product_category_id = product_categories.id").Where("product_brand_categories.deleted is null and product_brand_categories.product_brand_id = ?", id).Find(&productBrand.ProductCategories).Error; err != nil {
		return
	}
	return
}
func (repo *PostgresRepository) FindAll(params ...string) (dataCount int64, productBrands *[]dto.ProductBrand, err error) {
	if params[1] == "" {
		params[1] = "1"
	}
	if params[2] == "" {
		params[2] = "5"
	}
	var status bool
	if params[3] == "false" {
		status = false
	} else {
		status = true
	}
	nom, err := strconv.Atoi(params[1])
	if err != nil {
		return
	}
	if nom < 0 {
		params[1] = strconv.Itoa(nom)
	}
	den, err := strconv.Atoi(params[2])
	if err != nil {
		return
	}
	if den == -1 {
		if params[3] == "" {
			if err = repo.db.Debug().Table("product_brands").Select("product_brands.*, product_categories.name as category, count (products.id) as product").Joins("left join product_brand_categories on product_brand_categories.product_brand_id = product_brands.id").Joins("left join product_categories on product_categories.id = product_brand_categories.product_category_id").Joins("left join products on products.product_brand_category_id = product_brand_categories.id").Where("products.deleted is null and product_brands.deleted is null and (lower(product_brands.name) like lower(?) or lower(product_brands.id) like lower(?))", "%"+params[0]+"%", "%"+params[0]+"%").Group("product_brands.id,product_categories.name").Count(&dataCount).Scan(&productBrands).Error; err != nil {
				return
			}
			return
		}
		if err = repo.db.Debug().Table("product_brands").Select("product_brands.*, product_categories.name as category, count (products.id) as product").Joins("left join product_brand_categories on product_brand_categories.product_brand_id = product_brands.id").Joins("left join product_categories on product_categories.id = product_brand_categories.product_category_id").Joins("left join products on products.product_brand_category_id = product_brand_categories.id").Where("products.deleted is null and product_brands.deleted is null and (lower(product_brands.name) like lower(?) or lower(product_brands.id) like lower(?)) and product_brands.is_available = ?", "%"+params[0]+"%", "%"+params[0]+"%", &status).Group("product_brands.id,product_categories.name").Count(&dataCount).Scan(&productBrands).Error; err != nil {
			return
		}
		return
	}
	if params[3] == "" {
		if err = repo.db.Debug().Table("product_brands").Select("product_brands.*, product_categories.name as category, count (products.id) as product").Joins("left join product_brand_categories on product_brand_categories.product_brand_id = product_brands.id").Joins("left join product_categories on product_categories.id = product_brand_categories.product_category_id").Joins("left join products on products.product_brand_category_id = product_brand_categories.id").Where("products.deleted is null and product_brands.deleted is null and (lower(product_brands.name) like lower(?) or lower(product_brands.id) like lower(?))", "%"+params[0]+"%", "%"+params[0]+"%").Group("product_brands.id,product_categories.name").Count(&dataCount).Scopes(Paginate(params[1], params[2])).Scan(&productBrands).Error; err != nil {
			return
		}
		return
	}
	if err = repo.db.Debug().Table("product_brands").Select("product_brands.*, product_categories.name as category, count (products.id) as product").Joins("left join product_brand_categories on product_brand_categories.product_brand_id = product_brands.id").Joins("left join product_categories on product_categories.id = product_brand_categories.product_category_id").Joins("left join products on products.product_brand_category_id = product_brand_categories.id").Where("products.deleted is null and product_brands.deleted is null and (lower(product_brands.name) like lower(?) or lower(product_brands.id) like lower(?)) and product_brands.is_available = ?", "%"+params[0]+"%", "%"+params[0]+"%", &status).Group("product_brands.id,product_categories.name").Count(&dataCount).Scopes(Paginate(params[1], params[2])).Scan(&productBrands).Error; err != nil {
		return
	}
	return
}
func (repo *PostgresRepository) FindCategoryById(id string) (productCategory *models.ProductCategory, err error) {
	if err = repo.db.First(&productCategory, &id).Error; err != nil {
		return
	}
	return
}
func (repo *PostgresRepository) Insert(data *models.ProductBrand) (productBrand *models.ProductBrandResponse, err error) {
	if err = repo.db.Create(data).Preload("ProductCategories").First(&models.ProductBrand{}, &data.Id).Scan(&productBrand).Error; err != nil {
		return
	}
	return
}
func (repo *PostgresRepository) Update(id string, data *models.ProductBrand) (productBrand *models.ProductBrandResponse, err error) {
	log.Print(data)
	if err = repo.db.First(&productBrand, &id).Model(productBrand).Updates(&data).Error; err != nil {
		return
	}
	return
}
func (repo *PostgresRepository) Delete(id string) (err error) {
	if err = repo.db.Delete(&models.ProductBrand{}, &id).Error; err != nil {
		return
	}
	if err = repo.db.Delete(&models.ProductBrandCategory{}, "product_brand_id = ?", id).Error; err != nil {
		return
	}
	return
}
func (repo *PostgresRepository) CheckBrandCategory(brandId string, categoryId string) (rowCount int64, err error) {
	rowCount = repo.db.Where("deleted is null and product_brand_id = ? AND product_category_id = ?", brandId, categoryId).Find(&models.ProductBrandCategory{}).RowsAffected
	if rowCount > 0 {
		return rowCount, errors.New("")
	}
	return
}
func (repo *PostgresRepository) InsertBrandCategory(brandId string, categoryId string, slug string) (productBrand *models.ProductBrandCategory, err error) {
	productBrand = &models.ProductBrandCategory{
		Id:                uuid.New().String(),
		ProductBrandId:    brandId,
		ProductCategoryId: categoryId,
		Slug:              slug,
		IsAvailable:       nil,
	}
	if err = repo.db.Create(productBrand).Error; err != nil {
		return
	}
	return
}
func (repo *PostgresRepository) DeleteBrandCategory(brandId string, categoryId string) (err error) {
	if err = repo.db.Where("deleted is null and product_brand_id = ? AND product_category_id = ?", brandId, categoryId).Delete(&models.ProductBrandCategory{}).Error; err != nil {
		return
	}
	return
}
