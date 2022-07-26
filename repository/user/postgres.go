package user

import (
	"backend_capstone/models"
	"backend_capstone/services/user/dto"
	"log"
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
func (repo *PostgresRepository) DashboardDetailUser(params ...string) (data dto.UserDashboadDTO, err error) {
	if params[1] == "" {
		params[1] = "1"
	}
	if params[2] == "" {
		params[2] = "10"
	}
	den, err := strconv.Atoi(params[2])
	if err != nil {
		return
	}
	if den == -1 {
		if err = repo.db.Model(&data.Datas).Select("users.id as id,users.name as name, users.phone as phone, users.email as email,sum(payments.billed) as spending, count(transactions.id) as transactions").Joins("left join transactions on transactions.user_id = users.id left join payments on payments.transaction_id = transactions.id").Where("transactions.deleted is null and (lower(users.name) like lower(?) or lower(users.email) like lower(?) or lower(users.phone) like lower(?))", "%"+params[0]+"%", "%"+params[0]+"%", "%"+params[0]+"%").Group("users.id").Count(&data.PageLength).Find(&data.Datas).Error; err != nil {
			return
		}
		return
	}
	if err = repo.db.Model(&data.Datas).Select("users.id as id,users.name as name, users.phone as phone, users.email as email,sum(payments.billed) as spending, count(transactions.id) as transactions").Joins("left join transactions on transactions.user_id = users.id left join payments on payments.transaction_id = transactions.id").Where("transactions.deleted is null and (lower(users.name) like lower(?) or lower(users.email) like lower(?) or lower(users.phone) like lower(?))", "%"+params[0]+"%", "%"+params[0]+"%", "%"+params[0]+"%").Group("users.id").Count(&data.PageLength).Scopes(Paginate(params[1], params[2])).Find(&data.Datas).Error; err != nil {
		return
	}
	return
}
func (repo *PostgresRepository) FindById(id string) (user *models.UserResponse, err error) {
	log.Print(id)
	err = repo.db.Debug().First(&user, &id).Error
	if err != nil {
		return
	}
	return
}
func (repo *PostgresRepository) FindByIdentifier(identifier string) (user *models.User, err error) {
	err = repo.db.Debug().Where("email = ? or username = ?", identifier, identifier).First(&user).Error
	if err != nil {
		return
	}
	return
}
func (repo *PostgresRepository) FindAll() (users *[]models.UserResponse, err error) {
	err = repo.db.Find(&users).Error
	if err != nil {
		return
	}
	return
}
func (repo *PostgresRepository) Insert(data *models.User) (user *models.UserResponse, err error) {
	log.Print(repo.db)
	err = repo.db.Create(data).First(&user, &data.Id).Error
	if err != nil {
		log.Print("terjadi error")
		return
	}

	return
}
func (repo *PostgresRepository) Update(id string, data *models.User) (user *models.UserResponse, err error) {
	err = repo.db.First(&models.User{}, &id).Select("name", "phone", "password").UpdateColumns(&data).Error
	if err != nil {
		return
	}
	err = repo.db.First(&user, &id).Error
	if err != nil {
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
