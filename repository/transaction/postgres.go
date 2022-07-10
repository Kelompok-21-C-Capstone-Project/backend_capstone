package transaction

import (
	"backend_capstone/models"
	"backend_capstone/services/transaction/dto"
	"errors"
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

func (repo *PostgresRepository) FindById(id string) (transaction *models.Transaction, err error) {
	if err = repo.db.Preload("Payment").First(&transaction, &id).Error; err != nil {
		return
	}
	return
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

func (repo *PostgresRepository) UsersFindAll(uid string, params ...string) (transactions *[]dto.ClientTransactionsResponse, err error) {
	if params[1] != "" {
		if err = repo.db.Debug().Scopes(Paginate(params[5], params[6])).Order("transactions.created_at desc").Table("transactions").Select("transactions.id, product_categories.slug as category, product_categories.icon as icon, payments.status as status, products.name as product, transactions.description as transaction_details, payments.billed as charged, payments.created_at, payments.method as payment_method").Joins("left join payments on payments.transaction_id = transactions.id").Joins("left join products on transactions.product_id = products.id").Joins("left join product_brand_categories on products.product_brand_category_id = product_brand_categories.id").Joins("left join product_categories on product_categories.id = product_brand_categories.product_category_id").Where("transactions.user_id = ? and lower(product_categories.slug) like lower(?) and lower(payments.status) like lower(?) and to_char(transactions.created_at,'DD-MM-YYYY') = ? and (lower(transactions.id) like lower(?) or lower(products.name) like lower(?))", uid, params[4], params[3], params[1], params[0], params[0]).Scan(&transactions).Error; err != nil {
			return
		}
		return
	} else if params[2] != "" {
		if err = repo.db.Debug().Scopes(Paginate(params[5], params[6])).Order("transactions.created_at desc").Table("transactions").Select("transactions.id, product_categories.slug as category, product_categories.icon as icon, payments.status as status, products.name as product, transactions.description as transaction_details, payments.billed as charged, payments.created_at, payments.method as payment_method").Joins("left join payments on payments.transaction_id = transactions.id").Joins("left join products on transactions.product_id = products.id").Joins("left join product_brand_categories on products.product_brand_category_id = product_brand_categories.id").Joins("left join product_categories on product_categories.id = product_brand_categories.product_category_id").Where("transactions.user_id = ? and lower(product_categories.slug) like lower(?) and lower(payments.status) like lower(?) and (lower(transactions.id) like lower(?) or lower(products.name) like lower(?)) and transactions.created_at >= ? and transactions.created_at <= ?", uid, params[4], params[3], params[0], params[0], params[7], params[8]).Scan(&transactions).Error; err != nil {
			return
		}
		return
	}
	if err = repo.db.Debug().Scopes(Paginate(params[5], params[6])).Order("transactions.created_at desc").Table("transactions").Select("transactions.id, product_categories.slug as category, product_categories.icon as icon, payments.status as status, products.name as product, transactions.description as transaction_details, payments.billed as charged, payments.created_at, payments.method as payment_method").Joins("left join payments on payments.transaction_id = transactions.id").Joins("left join products on transactions.product_id = products.id").Joins("left join product_brand_categories on products.product_brand_category_id = product_brand_categories.id").Joins("left join product_categories on product_categories.id = product_brand_categories.product_category_id").Where("transactions.user_id = ? and lower(product_categories.slug) like lower(?) and lower(payments.status) like lower(?) and (lower(transactions.id) like lower(?) or lower(products.name) like lower(?))", uid, params[4], params[3], params[0], params[0]).Scan(&transactions).Error; err != nil {
		return
	}
	return
}
func (repo *PostgresRepository) UsersFindById(uid string, tid string) (transaction *dto.ClientTransactionsResponse, err error) {
	if err = repo.db.Order("created_at asc").Table("transactions").Select("transactions.id, product_categories.slug as category, product_categories.icon as icon,  payments.status as status, products.name as product, transactions.description as transaction_details, payments.billed as charged, payments.created_at, payments.method as payment_method").Joins("left join payments on payments.transaction_id = transactions.id").Joins("left join products on transactions.product_id = products.id").Joins("left join product_brand_categories on products.product_brand_category_id = product_brand_categories.id").Joins("left join product_categories on product_categories.id = product_brand_categories.product_category_id").Where("transactions.user_id = ? and transactions.id = ?", uid, tid).Scan(&transaction).Error; err != nil {
		return
	}
	return
}
func (repo *PostgresRepository) FindAll() (transactions *[]dto.ClientTransactionsResponse, err error) {
	if err = repo.db.Order("created_at desc").Table("transactions").Select("transactions.id, product_categories.slug as category, product_categories.icon as icon, payments.status as status, products.name as product, transactions.description as transaction_details, payments.billed as charged, payments.created_at, payments.method as payment_method").Joins("left join payments on payments.transaction_id = transactions.id").Joins("left join products on transactions.product_id = products.id").Joins("left join product_brand_categories on products.product_brand_category_id = product_brand_categories.id").Joins("left join product_categories on product_categories.id = product_brand_categories.product_category_id").Scan(&transactions).Error; err != nil {
		return
	}
	return
}
func (repo *PostgresRepository) Insert(data *models.Transaction) (transaction *models.Transaction, err error) {
	if err = repo.db.Create(&data).Error; err != nil {
		return
	}
	return data, err
}
func (repo *PostgresRepository) InsertPayment(data *models.Payment) (transaction *models.Payment, err error) {
	if err = repo.db.Create(&data).Error; err != nil {
		return
	}
	return data, err
}
func (repo *PostgresRepository) MidtransUpdate(tid string, status string) (err error) {
	if err = repo.db.First(&models.Payment{}, "transaction_id = ?", &tid).Update("status", status).Error; err != nil {
		return
	}
	return
}
func (repo *PostgresRepository) Update() (transaction *models.Transaction, err error) {
	return
}
func (repo *PostgresRepository) Delete(id string) (err error) {
	return
}

func (repo *PostgresRepository) CheckProductStock(pid string) (product *models.Product, err error) {
	data := new(models.Product)
	if err = repo.db.First(&data, &pid).Error; err != nil {
		return
	}
	if data.Stock < 1 {
		err = errors.New("barang habis")
		return
	} else {
		if err = repo.db.First(&data, &pid).Update("stock", data.Stock-1).Error; err != nil {
			return
		}
	}
	return data, err
}
func (repo *PostgresRepository) ProductReStock(pid string) (err error) {
	data := new(models.Product)
	if err = repo.db.First(&data, &pid).Update("stock", data.Stock+1).Error; err != nil {
		return
	}
	return err
}
func (repo *PostgresRepository) GetTransactionProduct(pid string) (product *models.Product, err error) {
	if err = repo.db.First(&product, &pid).Where("deleted is null").Error; err != nil {
		return
	}
	return
}
