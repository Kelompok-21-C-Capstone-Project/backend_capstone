package gormdriver

import (
	"backend_capstone/models"
	"backend_capstone/utils"
	"log"

	"gorm.io/gorm"
)

type GormMigrationService struct {
	gorm *gorm.DB
}

func NewGormMigrationService(dbCon *utils.DatabaseConnection) *GormMigrationService {
	log.Print("Enter NewGormMigrationService")
	switch dbCon.Driver {
	case utils.Postgres:
		return &GormMigrationService{
			gorm: dbCon.Postgres,
		}
	case utils.MySQL:
		return &GormMigrationService{
			gorm: dbCon.MySQL,
		}
	default:
		return nil
	}
}

func (s *GormMigrationService) GormMigrate() (err error) {
	log.Print("Enter GormMigration()")
	err = s.gorm.SetupJoinTable(&models.ProductBrand{}, "ProductCategories", &models.ProductBrandCategory{})
	if err != nil {
		return err
	}
	err = s.gorm.AutoMigrate(
		&models.ProductBrand{},
		&models.ProductCategory{},
		&models.ProductBrandCategory{},
		&models.Product{},
		&models.User{},
		&models.PaymentMethod{},
		&models.PaymentVendor{},
		&models.Payment{},
		&models.Transaction{},
	)
	if err != nil {
		return err
	}
	return
}