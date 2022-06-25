package models

import (
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
	err = s.gorm.SetupJoinTable(&ProductBrand{}, "ProductCategories", &ProductBrandCategory{})
	if err != nil {
		return err
	}
	err = s.gorm.AutoMigrate(
		&ProductBrand{},
		&ProductCategory{},
		&ProductBrandCategory{},
		&Product{},
		&User{},
		&PaymentMethod{},
		&PaymentVendor{},
		&Payment{},
		&Transaction{},
	)
	if err != nil {
		return err
	}
	return
}
