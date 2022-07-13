package productcategory

import (
	"backend_capstone/models"
	"backend_capstone/services/productcategory/dto"
	"strconv"

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

func (repo *PostgresRepository) FindById(id string) (productCategory *models.ProductCategoryResponse, err error) {
	if err = repo.db.First(&productCategory, &id).Error; err != nil {
		return
	}
	return
}
func (repo *PostgresRepository) FindAll(params ...string) (DataCount int64, productCategories *[]dto.ProductCategory, err error) {
	if params[1] == "" {
		params[1] = "1"
	}
	if params[2] == "" {
		params[2] = "5"
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
	if den <= 0 {
		den = 5
	}
	if err = repo.db.Table("product_categories").Select("product_categories.* , count(products.id) as product").Joins("left join product_brand_categories on product_brand_categories.product_category_id = product_categories.id").Joins("left join products on product_brand_categories.id = products.product_brand_category_id").Where("lower(product_categories.id) like lower(?) or lower(product_categories.name) like lower(?)", "%"+params[0]+"%", "%"+params[0]+"%").Group("product_categories.id").Count(&DataCount).Scopes(Paginate(params[1], params[2])).Scan(&productCategories).Error; err != nil {
		return
	}
	return
}
func (repo *PostgresRepository) Insert(data *models.ProductCategory) (productCategory *models.ProductCategoryResponse, err error) {
	if err = repo.db.Create(&data).First(&productCategory, &data.Id).Error; err != nil {
		return
	}
	return
}
func (repo *PostgresRepository) Update(id string, data *models.ProductCategory) (productCategory *models.ProductCategoryResponse, err error) {
	if err = repo.db.First(&productCategory, &id).Model(productCategory).Updates(data).Error; err != nil {
		return
	}
	return
}
func (repo *PostgresRepository) Delete(id string) (err error) {
	if err = repo.db.Delete(&models.ProductCategory{}, &id).Error; err != nil {
		return
	}
	return
}
