package transaction

import (
	"backend_capstone/models"
	"backend_capstone/services/transaction/dto"
	"errors"

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
func (repo *PostgresRepository) FindByQuery(key string, value interface{}) (transactions *[]models.Transaction, err error) {
	return
}

func (repo *PostgresRepository) UsersFindAll(uid string) (transactions *[]dto.ClientTransactionsResponse, err error) {
	if err = repo.db.Order("created_at desc").Table("transactions").Select("transactions.id, product_categories.name as category, product_categories.slug as icon, payments.status as status, products.name as product, transactions.description as transaction_details, payments.billed as charged, payments.created_at, payments.method as payment_method").Joins("left join payments on payments.transaction_id = transactions.id").Joins("left join products on transactions.product_id = products.id").Joins("left join product_brand_categories on products.product_brand_category_id = product_brand_categories.id").Joins("left join product_categories on product_categories.id = product_brand_categories.product_category_id").Where("transactions.user_id = ?", uid).Scan(&transactions).Error; err != nil {
		return
	}
	return
}
func (repo *PostgresRepository) UsersFindById(uid string, tid string) (transaction *dto.ClientTransactionsResponse, err error) {
	if err = repo.db.Order("created_at asc").Table("transactions").Select("transactions.id, product_categories.name as category, product_categories.slug as icon,  payments.status as status, products.name as product, transactions.description as transaction_details, payments.billed as charged, payments.created_at, payments.method as payment_method").Joins("left join payments on payments.transaction_id = transactions.id").Joins("left join products on transactions.product_id = products.id").Joins("left join product_brand_categories on products.product_brand_category_id = product_brand_categories.id").Joins("left join product_categories on product_categories.id = product_brand_categories.product_category_id").Where("transactions.user_id = ? and transactions.id = ?", uid, tid).Scan(&transaction).Error; err != nil {
		return
	}
	return
}
func (repo *PostgresRepository) FindAll() (transactions *[]dto.ClientTransactionsResponse, err error) {
	if err = repo.db.Order("created_at asc").Table("transactions").Select("transactions.id, product_categories.name as category, product_categories.slug as icon, payments.status as status, products.name as product, transactions.description as transaction_details, payments.billed as charged, payments.created_at, payments.method as payment_method").Joins("left join payments on payments.transaction_id = transactions.id").Joins("left join products on transactions.product_id = products.id").Joins("left join product_brand_categories on products.product_brand_category_id = product_brand_categories.id").Joins("left join product_categories on product_categories.id = product_brand_categories.product_category_id").Scan(&transactions).Error; err != nil {
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
