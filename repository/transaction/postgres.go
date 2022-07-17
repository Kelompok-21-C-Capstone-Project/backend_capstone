package transaction

import (
	"backend_capstone/models"
	"backend_capstone/services/transaction/dto"
	"errors"
	"log"
	"strconv"
	"time"

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

func (repo *PostgresRepository) CountUsers(id string) (ammount int64, err error) {
	// repo.db.Model(&models.Transaction).Where("deleted is null and user_id = ?", id).C
	return
}

func (repo *PostgresRepository) FindById(id string) (transaction *models.TransactionResponse, err error) {
	if err = repo.db.Preload("Payment").First(&models.Transaction{}, &id).Scan(&transaction).Error; err != nil {
		return
	}
	if err = repo.db.First(&models.Payment{}, "transaction_id = ?", id).Scan(&transaction.Payment).Error; err != nil {
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

func (repo *PostgresRepository) UsersFindAll(uid string, params ...string) (dataCount int64, transactions *[]dto.ClientTransactionsResponse, err error) {
	den, err := strconv.Atoi(params[6])
	if err != nil {
		return
	}
	if params[1] != "" {
		if den == -1 {
			if err = repo.db.Debug().Order("transactions.created_at desc").Table("transactions").Select("transactions.id, product_categories.slug as category, product_categories.icon as icon, payments.status as status, products.id as product_id, products.name as product, transactions.description as transaction_details, payments.billed as charged, payments.created_at::timestamp at time zone 'gmt+7' as created_at, payments.method as payment_method").Joins("left join payments on payments.transaction_id = transactions.id").Joins("left join products on transactions.product_id = products.id").Joins("left join product_brand_categories on products.product_brand_category_id = product_brand_categories.id").Joins("left join product_categories on product_categories.id = product_brand_categories.product_category_id").Where("transactions.deleted is null and transactions.user_id = ? and lower(product_categories.slug) like lower(?) and lower(payments.status) like lower(?) and to_char(transactions.created_at::timestamp at time zone 'gmt+7','DD-MM-YYYY') = ? and (lower(transactions.id) like lower(?) or lower(products.name) like lower(?))", uid, params[4], params[3], params[1], params[0], params[0]).Count(&dataCount).Scan(&transactions).Error; err != nil {
				return
			}
			return
		}
		if err = repo.db.Debug().Order("transactions.created_at desc").Table("transactions").Select("transactions.id, product_categories.slug as category, product_categories.icon as icon, payments.status as status, products.id as product_id, products.name as product, transactions.description as transaction_details, payments.billed as charged, payments.created_at::timestamp at time zone 'gmt+7' as created_at, payments.method as payment_method").Joins("left join payments on payments.transaction_id = transactions.id").Joins("left join products on transactions.product_id = products.id").Joins("left join product_brand_categories on products.product_brand_category_id = product_brand_categories.id").Joins("left join product_categories on product_categories.id = product_brand_categories.product_category_id").Where("transactions.deleted is null and transactions.user_id = ? and lower(product_categories.slug) like lower(?) and lower(payments.status) like lower(?) and to_char(transactions.created_at::timestamp at time zone 'gmt+7','DD-MM-YYYY') = ? and (lower(transactions.id) like lower(?) or lower(products.name) like lower(?))", uid, params[4], params[3], params[1], params[0], params[0]).Count(&dataCount).Scopes(Paginate(params[5], params[6])).Scan(&transactions).Error; err != nil {
			return
		}
		return
	} else if params[2] != "" {
		if den == -1 {
			if err = repo.db.Debug().Order("transactions.created_at desc").Table("transactions").Select("transactions.id, product_categories.slug as category, product_categories.icon as icon, payments.status as status, products.id as product_id, products.name as product, transactions.description as transaction_details, payments.billed as charged, payments.created_at::timestamp at time zone 'gmt+7' as created_at, payments.method as payment_method").Joins("left join payments on payments.transaction_id = transactions.id").Joins("left join products on transactions.product_id = products.id").Joins("left join product_brand_categories on products.product_brand_category_id = product_brand_categories.id").Joins("left join product_categories on product_categories.id = product_brand_categories.product_category_id").Where("transactions.deleted is null and transactions.user_id = ? and lower(product_categories.slug) like lower(?) and lower(payments.status) like lower(?) and (lower(transactions.id) like lower(?) or lower(products.name) like lower(?)) and transactions.created_at >= ? and transactions.created_at <= ?", uid, params[4], params[3], params[0], params[0], params[7], params[8]).Count(&dataCount).Scan(&transactions).Error; err != nil {
				return
			}
			return
		}
		if err = repo.db.Debug().Order("transactions.created_at desc").Table("transactions").Select("transactions.id, product_categories.slug as category, product_categories.icon as icon, payments.status as status, products.id as product_id, products.name as product, transactions.description as transaction_details, payments.billed as charged, payments.created_at::timestamp at time zone 'gmt+7' as created_at, payments.method as payment_method").Joins("left join payments on payments.transaction_id = transactions.id").Joins("left join products on transactions.product_id = products.id").Joins("left join product_brand_categories on products.product_brand_category_id = product_brand_categories.id").Joins("left join product_categories on product_categories.id = product_brand_categories.product_category_id").Where("transactions.deleted is null and transactions.user_id = ? and lower(product_categories.slug) like lower(?) and lower(payments.status) like lower(?) and (lower(transactions.id) like lower(?) or lower(products.name) like lower(?)) and transactions.created_at >= ? and transactions.created_at <= ?", uid, params[4], params[3], params[0], params[0], params[7], params[8]).Count(&dataCount).Scopes(Paginate(params[5], params[6])).Scan(&transactions).Error; err != nil {
			return
		}
		return
	}
	if den == -1 {
		if err = repo.db.Debug().Order("transactions.created_at desc").Table("transactions").Select("transactions.id, product_categories.slug as category, product_categories.icon as icon, payments.status as status, products.id as product_id, products.name as product, transactions.description as transaction_details, payments.billed as charged, payments.created_at::timestamp at time zone 'gmt+7' as created_at, payments.method as payment_method").Joins("left join payments on payments.transaction_id = transactions.id").Joins("left join products on transactions.product_id = products.id").Joins("left join product_brand_categories on products.product_brand_category_id = product_brand_categories.id").Joins("left join product_categories on product_categories.id = product_brand_categories.product_category_id").Where("transactions.deleted is null and transactions.user_id = ? and lower(product_categories.slug) like lower(?) and lower(payments.status) like lower(?) and (lower(transactions.id) like lower(?) or lower(products.name) like lower(?))", uid, params[4], params[3], params[0], params[0]).Count(&dataCount).Scan(&transactions).Error; err != nil {
			return
		}
		return
	}
	if err = repo.db.Debug().Order("transactions.created_at desc").Table("transactions").Select("transactions.id, product_categories.slug as category, product_categories.icon as icon, payments.status as status, products.id as product_id, products.name as product, transactions.description as transaction_details, payments.billed as charged, payments.created_at::timestamp at time zone 'gmt+7' as created_at, payments.method as payment_method").Joins("left join payments on payments.transaction_id = transactions.id").Joins("left join products on transactions.product_id = products.id").Joins("left join product_brand_categories on products.product_brand_category_id = product_brand_categories.id").Joins("left join product_categories on product_categories.id = product_brand_categories.product_category_id").Where("transactions.deleted is null and transactions.user_id = ? and lower(product_categories.slug) like lower(?) and lower(payments.status) like lower(?) and (lower(transactions.id) like lower(?) or lower(products.name) like lower(?))", uid, params[4], params[3], params[0], params[0]).Count(&dataCount).Scopes(Paginate(params[5], params[6])).Scan(&transactions).Error; err != nil {
		return
	}
	return
}
func (repo *PostgresRepository) UsersFindById(uid string, tid string) (transaction *dto.ClientTransactionsResponse, err error) {
	if err = repo.db.Order("created_at asc").Table("transactions").Select("transactions.id, product_categories.slug as category, product_categories.icon as icon,  payments.status as status, products.name as product, transactions.description as transaction_details, payments.billed as charged, payments.created_at::timestamp at time zone 'gmt+7' as created_at, payments.method as payment_method").Joins("left join payments on payments.transaction_id = transactions.id").Joins("left join products on transactions.product_id = products.id").Joins("left join product_brand_categories on products.product_brand_category_id = product_brand_categories.id").Joins("left join product_categories on product_categories.id = product_brand_categories.product_category_id").Where("transactions.deleted is null and transactions.user_id = ? and transactions.id = ?", uid, tid).Scan(&transaction).Error; err != nil {
		return
	}
	return
}
func (repo *PostgresRepository) FindAll() (transactions *[]dto.ClientTransactionsResponse, err error) {
	if err = repo.db.Order("created_at desc").Table("transactions").Select("transactions.id, product_categories.slug as category, product_categories.icon as icon, payments.status as status, products.id as product_id, products.name as product, transactions.description as transaction_details, payments.billed as charged, payments.created_at::timestamp at time zone 'gmt+7' as created_at, payments.method as payment_method").Joins("left join payments on payments.transaction_id = transactions.id").Joins("left join products on transactions.product_id = products.id").Joins("left join product_brand_categories on products.product_brand_category_id = product_brand_categories.id").Joins("left join product_categories on product_categories.id = product_brand_categories.product_category_id").Scan(&transactions).Error; err != nil {
		return
	}
	return
}
func (repo *PostgresRepository) Insert(data *models.Transaction) (transaction *models.TransactionResponse, err error) {
	if err = repo.db.Create(&data).First(&data, &data.Id).Scan(&transaction).Error; err != nil {
		return
	}
	return
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
func (repo *PostgresRepository) Update() (transaction *models.TransactionResponse, err error) {
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
func (repo *PostgresRepository) GetBillById(tid string) (bill dto.BillClient, err error) {
	if err = repo.db.Debug().Table("transactions").Select("payments.id as id, transactions.id as transaction_id, payments.status as status, payments.description as va_number, payments.method_details as payment_details, payments.billed as billed, payments.charged as charger, products.name as product, products.price as product_price, payments.created_at as deadline").Joins("left join payments on payments.transaction_id = transactions.id").Joins("left join products on products.id = transactions.product_id").Where("transactions.deleted is null and transactions", &tid).Scan(&bill).Error; err != nil {
		return
	}
	return
}
func (repo *PostgresRepository) GetUserInfo(tid string) (user models.UserResponse, err error) {
	var transaction models.Transaction
	if err = repo.db.First(&transaction, &tid).Error; err != nil {
		return
	}
	if err = repo.db.First(&models.User{}, &transaction.UserId).Scan(&user).Error; err != nil {
		log.Print("terjadi error")
		return
	}
	return
}
func (repo *PostgresRepository) AdminDetailTransaction(params ...string) (dashboardData dto.DashboardDetailTransactionDTO, err error) {
	log.Print("PostgresRepository.AdminDetailTransaction")
	if params[5] == "" {
		params[5] = "10"
	}
	den, err := strconv.Atoi(params[5])
	if err != nil {
		return
	}
	if params[6] == "" {
		params[6] = time.Now().Format("02-01-2006")
	}
	if params[7] == "" {
		params[7] = time.Now().AddDate(0, 0, 1).Format("02-01-2006")
	}
	if err = repo.db.Debug().Table("transactions").Select("sum(payments.billed) as sum, sum(payments.charged) as profit, count(transactions.id) as transaction_count").Joins("left join products on products.id = transactions.product_id left join payments on transactions.id = payments.transaction_id left join product_brand_categories on product_brand_categories.id = products.product_brand_category_id left join product_brands on product_brand_categories.product_brand_id = product_brands.id left join product_categories on product_brand_categories.product_category_id = product_categories.id").Where("transactions.deleted is null and products.deleted is null and lower(product_categories.slug) like lower(?) and (lower(product_brands.name) like lower(?) or lower(payments.method_details) like lower(?) or lower(products.name) like lower(?)) and transactions.created_at >= ? and transactions.created_at <= ? and lower(status) like lower(?)", "%"+params[3]+"%", "%"+params[0]+"%", "%"+params[0]+"%", "%"+params[0]+"%", params[6], params[7], "%"+params[2]+"%").Scan(&dashboardData.Sum).Error; err != nil {
		return
	}
	if den == -1 {
		if err = repo.db.Debug().Table("transactions").Select("payments.id as id, payments.status as status, to_char(transactions.created_at::timestamp at time zone 'gmt+7','DD/MM/YYYY HH24:MI:SS') as order_time, to_char(payments.updated_at::timestamp at time zone 'gmt+7','DD/MM/YYYY HH24:MI:SS') as payment_time, product_categories.name as category, users.id as user_id, users.name as user_name, payments.billed as paid, payments.method_details as payment_method").Joins("left join products on products.id = transactions.product_id left join payments on transactions.id = payments.transaction_id left join product_brand_categories on product_brand_categories.id = products.product_brand_category_id left join product_brands on product_brand_categories.product_brand_id = product_brands.id left join product_categories on product_brand_categories.product_category_id = product_categories.id left join users on users.id = transactions.user_id").Where("transactions.deleted is null and products.deleted is null and lower(product_categories.slug) like lower(?) and (lower(product_brands.name) like lower(?) or lower(payments.method_details) like lower(?) or lower(products.name) like lower(?)) and transactions.created_at >= ? and transactions.created_at <= ? and lower(status) like lower(?)", "%"+params[3]+"%", "%"+params[0]+"%", "%"+params[0]+"%", "%"+params[0]+"%", params[6], params[7], "%"+params[2]+"%").Order("transactions.created_at asc").Count(&dashboardData.Transactions.PageLength).Scan(&dashboardData.Transactions.Data).Error; err != nil {
			return
		}
		return
	}
	if err = repo.db.Debug().Table("transactions").Select("payments.id as id, payments.status as status, to_char(transactions.created_at::timestamp at time zone 'gmt+7','DD/MM/YYYY HH24:MI:SS') as order_time, to_char(payments.updated_at::timestamp at time zone 'gmt+7','DD/MM/YYYY HH24:MI:SS') as payment_time, product_categories.name as category, users.id as user_id, users.name as user_name, payments.billed as paid, payments.method_details as payment_method").Joins("left join products on products.id = transactions.product_id left join payments on transactions.id = payments.transaction_id left join product_brand_categories on product_brand_categories.id = products.product_brand_category_id left join product_brands on product_brand_categories.product_brand_id = product_brands.id left join product_categories on product_brand_categories.product_category_id = product_categories.id left join users on users.id = transactions.user_id").Where("transactions.deleted is null and products.deleted is null and lower(product_categories.slug) like lower(?) and (lower(product_brands.name) like lower(?) or lower(payments.method_details) like lower(?) or lower(products.name) like lower(?)) and transactions.created_at >= ? and transactions.created_at <= ? and lower(status) like lower(?)", "%"+params[3]+"%", "%"+params[0]+"%", "%"+params[0]+"%", "%"+params[0]+"%", params[6], params[7], "%"+params[2]+"%").Order("transactions.created_at asc").Count(&dashboardData.Transactions.PageLength).Scopes(Paginate(params[4], params[5])).Scan(&dashboardData.Transactions.Data).Error; err != nil {
		return
	}
	return
}
