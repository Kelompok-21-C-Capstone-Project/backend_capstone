package user

import (
	"backend_capstone/models"
	"log"

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

func (repo *PostgresRepository) FindById(id string) (user *models.User, err error) {
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
func (repo *PostgresRepository) FindAll() (users *[]models.User, err error) {
	err = repo.db.Find(&users).Error
	if err != nil {
		return
	}
	return
}
func (repo *PostgresRepository) Insert(data *models.User) (user *models.User, err error) {
	err = repo.db.Create(data).Error
	if err != nil {
		return
	}
	return data, err
}
func (repo *PostgresRepository) Update(id string, data *models.User) (user *models.User, err error) {
	err = repo.db.First(user, &id).Select("name", "phone", "password").UpdateColumns(data).Error
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
