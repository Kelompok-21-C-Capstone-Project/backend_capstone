package products

import (
	"backend_capstone/models"
	"backend_capstone/services/product/dto"
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

func ProductTransaction() func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Table("transactions")
	}
}

func (repo *PostgresRepository) FindById(id string) (product *models.ProductResponse, err error) {
	if err = repo.db.Debug().Preload("ProductBrandCategory").Preload("Transactions").Preload("Transactions.Payment").First(&product, &id).Error; err != nil {
		return
	}
	return
}
func (repo *PostgresRepository) FindAll(params ...string) (dataCount int64, products *[]dto.Product, err error) {
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
	if params[3] != "" && params[4] != "" {
		if den == -1 {
			if err = repo.db.Debug().Table("products").Select("products.*, product_categories.name as category , product_brands.name as brand").Joins("left join product_brand_categories on product_brand_categories.id = products.product_brand_category_id").Joins("left join product_brands on product_brands.id = product_brand_categories.product_brand_id").Joins("left join product_categories on product_categories.id = product_brand_categories.product_category_id").Where("products.deleted is null and lower(product_brands.name) like lower(?) and lower(product_categories.slug) like lower(?) and (lower(products.id) like lower(?) or lower(products.name) like lower(?))", "%"+params[3]+"%", "%"+params[4]+"%", "%"+params[0]+"%", "%"+params[0]+"%").Count(&dataCount).Scan(&products).Error; err != nil {
				return
			}
			return
		}
		if err = repo.db.Debug().Table("products").Select("products.*, product_categories.name as category , product_brands.name as brand").Joins("left join product_brand_categories on product_brand_categories.id = products.product_brand_category_id").Joins("left join product_brands on product_brands.id = product_brand_categories.product_brand_id").Joins("left join product_categories on product_categories.id = product_brand_categories.product_category_id").Where("products.deleted is null and lower(product_brands.name) like lower(?) and lower(product_categories.slug) like lower(?) and (lower(products.id) like lower(?) or lower(products.name) like lower(?))", "%"+params[3]+"%", "%"+params[4]+"%", "%"+params[0]+"%", "%"+params[0]+"%").Count(&dataCount).Scopes(Paginate(params[1], params[2])).Scan(&products).Error; err != nil {
			return
		}
		return
	}
	if params[3] != "" {
		if den == -1 {
			if err = repo.db.Debug().Table("products").Select("products.*, product_categories.name as category , product_brands.name as brand").Joins("left join product_brand_categories on product_brand_categories.id = products.product_brand_category_id").Joins("left join product_brands on product_brands.id = product_brand_categories.product_brand_id").Joins("left join product_categories on product_categories.id = product_brand_categories.product_category_id").Where("products.deleted is null and lower(product_brands.name) like lower(?) and (lower(products.id) like lower(?) or lower(products.name) like lower(?))", "%"+params[3]+"%", "%"+params[0]+"%", "%"+params[0]+"%").Count(&dataCount).Scan(&products).Error; err != nil {
				return
			}
			return
		}
		if err = repo.db.Debug().Table("products").Select("products.*, product_categories.name as category , product_brands.name as brand").Joins("left join product_brand_categories on product_brand_categories.id = products.product_brand_category_id").Joins("left join product_brands on product_brands.id = product_brand_categories.product_brand_id").Joins("left join product_categories on product_categories.id = product_brand_categories.product_category_id").Where("products.deleted is null and lower(product_brands.name) like lower(?) and (lower(products.id) like lower(?) or lower(products.name) like lower(?))", "%"+params[3]+"%", "%"+params[0]+"%", "%"+params[0]+"%").Count(&dataCount).Scopes(Paginate(params[1], params[2])).Scan(&products).Error; err != nil {
			return
		}
		return
	}
	if params[4] != "" {
		if den == -1 {
			if err = repo.db.Debug().Table("products").Select("products.*, product_categories.name as category , product_brands.name as brand").Joins("left join product_brand_categories on product_brand_categories.id = products.product_brand_category_id").Joins("left join product_brands on product_brands.id = product_brand_categories.product_brand_id").Joins("left join product_categories on product_categories.id = product_brand_categories.product_category_id").Where("products.deleted is null and lower(product_categories.slug) like lower(?) and (lower(products.id) like lower(?) or lower(products.name) like lower(?))", "%"+params[4]+"%", "%"+params[0]+"%", "%"+params[0]+"%").Count(&dataCount).Scan(&products).Error; err != nil {
				return
			}
			return
		}
		if err = repo.db.Debug().Table("products").Select("products.*, product_categories.name as category , product_brands.name as brand").Joins("left join product_brand_categories on product_brand_categories.id = products.product_brand_category_id").Joins("left join product_brands on product_brands.id = product_brand_categories.product_brand_id").Joins("left join product_categories on product_categories.id = product_brand_categories.product_category_id").Where("products.deleted is null and lower(product_categories.slug) like lower(?) and (lower(products.id) like lower(?) or lower(products.name) like lower(?))", "%"+params[4]+"%", "%"+params[0]+"%", "%"+params[0]+"%").Count(&dataCount).Scopes(Paginate(params[1], params[2])).Scan(&products).Error; err != nil {
			return
		}
		return
	}
	if den == -1 {
		if err = repo.db.Debug().Table("products").Select("products.*, product_categories.name as category , product_brands.name as brand").Joins("left join product_brand_categories on product_brand_categories.id = products.product_brand_category_id").Joins("left join product_brands on product_brands.id = product_brand_categories.product_brand_id").Joins("left join product_categories on product_categories.id = product_brand_categories.product_category_id").Where("products.deleted is null and (lower(products.id) like lower(?) or lower(products.name) like lower(?))", "%"+params[0]+"%", "%"+params[0]+"%").Count(&dataCount).Scan(&products).Error; err != nil {
			return
		}
		return
	}
	if err = repo.db.Debug().Table("products").Select("products.*, product_categories.name as category , product_brands.name as brand").Joins("left join product_brand_categories on product_brand_categories.id = products.product_brand_category_id").Joins("left join product_brands on product_brands.id = product_brand_categories.product_brand_id").Joins("left join product_categories on product_categories.id = product_brand_categories.product_category_id").Where("products.deleted is null and (lower(products.id) like lower(?) or lower(products.name) like lower(?))", "%"+params[0]+"%", "%"+params[0]+"%").Count(&dataCount).Scopes(Paginate(params[1], params[2])).Scan(&products).Error; err != nil {
		return
	}
	return
}
func (repo *PostgresRepository) ClientFindAllBySlug(slug string) (products *dto.ProductCategory, err error) {
	err = repo.db.Table("product_categories").Select("product_categories.id, product_categories.name, product_categories.slug").Find(&products, "slug = ?", &slug).Error
	if err != nil {
		return
	}
	if err = repo.db.Table("product_brand_categories").Select("products.id, product_brands.name as group, products.name as label, products.description as description, products.stock as stock, products.price as price, products.is_discount").Joins("left join product_categories on product_brand_categories.product_category_id = product_categories.id").Joins("left join products on product_brand_categories.id = products.product_brand_category_id").Joins("left join product_brands on product_brand_categories.product_brand_id = product_brands.id").Where("product_brand_categories.deleted is null and product_brand_categories.product_category_id = ? and products.deleted is null", products.Id).Find(&products.Products).Error; err != nil {
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
		if err = repo.db.Table("product_brand_categories").Select("products.id, product_brands.name as group, products.name as label, products.description as description, products.stock as stock, products.price as price, products.is_discount").Joins("left join product_categories on product_brand_categories.product_category_id = product_categories.id").Joins("left join products on product_brand_categories.id = products.product_brand_category_id").Joins("left join product_brands on product_brand_categories.product_brand_id = product_brands.id").Where("product_brand_categories.deleted is null and product_brand_categories.product_category_id = ?", el.Id).Find(&(*products)[i].Products).Error; err != nil {
			return
		}
	}
	return
}
func (repo *PostgresRepository) Insert(data *models.Product) (product *models.ProductResponse, err error) {
	if err = repo.db.Create(&data).First(&product, &data.Id).Error; err != nil {
		return
	}
	return product, err
}
func (repo *PostgresRepository) Update(id string, data *models.Product) (product *models.ProductResponse, err error) {
	if err = repo.db.First(&product, &id).Model(product).Updates(data).Preload("ProductBrandCategory").First(&models.Product{}, &id).Scan(&product).Error; err != nil {
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
	rowCount := repo.db.Where("deleted is null and product_brand_id = ? AND product_category_id = ?", brandId, categoryId).Find(&check).RowsAffected
	if rowCount > 0 {
		return check.Id, err
	}
	id := uuid.New()
	if err = repo.db.Create(&models.ProductBrandCategory{
		Id:                id.String(),
		ProductBrandId:    brandId,
		ProductCategoryId: categoryId,
		Slug:              "",
		IsAvailable:       nil,
	}).Error; err != nil {
		return
	}
	return id.String(), err
}
func (repo *PostgresRepository) UpdateStock(data *dto.UpdateStockDTO) (err error) {
	for _, el := range (*data).Datas {
		if err = repo.db.First(&models.Product{}, &el.Id).Update("stock", gorm.Expr("stock + ?", el.Stock)).Error; err != nil {
			return
		}
	}
	return
}
