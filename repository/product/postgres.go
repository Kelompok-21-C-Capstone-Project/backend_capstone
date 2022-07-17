package products

import (
	"backend_capstone/models"
	"backend_capstone/services/product/dto"
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
	var status bool
	if params[5] == "false" {
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
	// if params[3] != "" || params[4] != "" {
	if den == -1 {
		if params[5] == "" {
			if err = repo.db.Debug().Model(&products).Preload("Category").Preload("Brand").Select("products.*,product_brands.id as brand_id, product_categories.id as category_id").Joins("left join product_brand_categories on product_brand_categories.id = products.product_brand_category_id").Joins("left join product_brands on product_brands.id = product_brand_categories.product_brand_id").Joins("left join product_categories on product_categories.id = product_brand_categories.product_category_id").Where("products.deleted is null and lower(product_brands.name) like lower(?) and lower(product_categories.slug) like lower(?) and (lower(products.id) like lower(?) or lower(products.name) like lower(?))", "%"+params[3]+"%", "%"+params[4]+"%", "%"+params[0]+"%", "%"+params[0]+"%").Count(&dataCount).Find(&products).Error; err != nil {
				return
			}
			return
		}
		if err = repo.db.Debug().Model(&products).Preload("Category").Preload("Brand").Select("products.*,product_brands.id as brand_id, product_categories.id as category_id").Joins("left join product_brand_categories on product_brand_categories.id = products.product_brand_category_id").Joins("left join product_brands on product_brands.id = product_brand_categories.product_brand_id").Joins("left join product_categories on product_categories.id = product_brand_categories.product_category_id").Where("products.deleted is null and lower(product_brands.name) like lower(?) and lower(product_categories.slug) like lower(?) and (lower(products.id) like lower(?) or lower(products.name) like lower(?)) and products.is_available = ?", "%"+params[3]+"%", "%"+params[4]+"%", "%"+params[0]+"%", "%"+params[0]+"%", status).Count(&dataCount).Find(&products).Error; err != nil {
			return
		}
		return
	}
	if params[5] == "" {
		if err = repo.db.Debug().Model(&products).Preload("Category").Preload("Brand").Select("products.*,product_brands.id as brand_id, product_categories.id as category_id").Joins("left join product_brand_categories on product_brand_categories.id = products.product_brand_category_id").Joins("left join product_brands on product_brands.id = product_brand_categories.product_brand_id").Joins("left join product_categories on product_categories.id = product_brand_categories.product_category_id").Where("products.deleted is null and lower(product_brands.name) like lower(?) and lower(product_categories.slug) like lower(?) and (lower(products.id) like lower(?) or lower(products.name) like lower(?))", "%"+params[3]+"%", "%"+params[4]+"%", "%"+params[0]+"%", "%"+params[0]+"%").Count(&dataCount).Scopes(Paginate(params[1], params[2])).Find(&products).Error; err != nil {
			return
		}
		return
	}
	if err = repo.db.Debug().Model(&products).Preload("Category").Preload("Brand").Select("products.*,product_brands.id as brand_id, product_categories.id as category_id").Joins("left join product_brand_categories on product_brand_categories.id = products.product_brand_category_id").Joins("left join product_brands on product_brands.id = product_brand_categories.product_brand_id").Joins("left join product_categories on product_categories.id = product_brand_categories.product_category_id").Where("products.deleted is null and lower(product_brands.name) like lower(?) and lower(product_categories.slug) like lower(?) and (lower(products.id) like lower(?) or lower(products.name) like lower(?)) and products.is_available = ?", "%"+params[3]+"%", "%"+params[4]+"%", "%"+params[0]+"%", "%"+params[0]+"%", status).Count(&dataCount).Scopes(Paginate(params[1], params[2])).Find(&products).Error; err != nil {
		return
	}
	return
}
func (repo *PostgresRepository) ClientFindAllBySlug(slug string) (products *dto.ProductCategory, err error) {
	err = repo.db.Table("product_categories").Select("product_categories.id, product_categories.name, product_categories.slug").Find(&products, "slug = ?", &slug).Error
	if err != nil {
		return
	}
	if err = repo.db.Table("product_brand_categories").Select("products.id, product_brands.name as group, products.name as label, products.description as description, products.stock as stock, products.price as price, products.is_discount").Joins("left join product_categories on product_brand_categories.product_category_id = product_categories.id").Joins("left join products on product_brand_categories.id = products.product_brand_category_id").Joins("left join product_brands on product_brand_categories.product_brand_id = product_brands.id").Where("product_brand_categories.deleted is null and product_brand_categories.product_category_id = ? and products.deleted is null and products.is_available is true", products.Id).Find(&products.Products).Error; err != nil {
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
		if err = repo.db.Table("product_brand_categories").Select("products.id, product_brands.name as group, products.name as label, products.description as description, products.stock as stock, products.price as price, products.is_discount").Joins("left join product_categories on product_brand_categories.product_category_id = product_categories.id").Joins("left join products on product_brand_categories.id = products.product_brand_category_id").Joins("left join product_brands on product_brand_categories.product_brand_id = product_brands.id").Where("product_brand_categories.deleted is null and products.is_available is true and product_brand_categories.product_category_id = ?", el.Id).Scan(&(*products)[i].Products).Error; err != nil {
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
	if err = repo.db.First(&product, &id).Model(product).Update("Stock", gorm.Expr("stock + ?", data.Stock)).Omit("Stock").Updates(data).Preload("ProductBrandCategory").First(&models.Product{}, &id).Scan(&product).Error; err != nil {
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
			repo.db.Rollback()
			return
		}
	}
	return
}
func (repo *PostgresRepository) CreateSupplyProduct(dataSupply models.Supply, dataProducts []models.SupplyProduct) (err error) {
	if err = repo.db.Create(&dataSupply).Error; err != nil {
		return
	}
	if err = repo.db.Create(&dataProducts).Error; err != nil {
		return
	}
	return
}
func (repo *PostgresRepository) GetSupplyInvocie(params ...string) (data dto.DataSupplyDTO, err error) {
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
	log.Print(params[4], " ", params[5])
	if den == -1 {
		if err = repo.db.Model(&data.Data).Select("code_no as id, name, to_char(created_at::timestamp at time zone'gmt+7','DD-MM-YYYY') as date,sum_stock as sum)").Count(&data.PageLength).Where("(code_no like ? or lower(name) like lower(?)) and to_char(supplies.created_at::timestamp at time zone 'gmt+7','DD-MM-YYYY HH24:MI:SS') >= ? and to_char(supplies.created_at::timestamp at time zone 'gmt+7','DD-MM-YYYY HH24:MI:SS') < ?", "%"+params[0]+"%", "%"+params[0]+"%", params[4], params[5]).Find(&data.Data).Error; err != nil {
			return
		}
	}
	if err = repo.db.Debug().Model(&data.Data).Select("code_no as id, name, to_char(created_at::timestamp at time zone'gmt+7','DD-MM-YYYY') as date,sum_stock as sum").Count(&data.PageLength).Where("(code_no like ? or lower(name) like lower(?)) and to_char(supplies.created_at::timestamp at time zone 'gmt+7','DD-MM-YYYY HH24:MI:SS') >= ? and to_char(supplies.created_at::timestamp at time zone 'gmt+7','DD-MM-YYYY HH24:MI:SS') < ?", "%"+params[0]+"%", "%"+params[0]+"%", params[4], params[5]).Scopes(Paginate(params[1], params[2])).Find(&data.Data).Error; err != nil {
		return
	}
	return
}
