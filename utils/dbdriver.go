package utils

import (
	"backend_capstone/configs"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseDriver string

const (
	Postgres DatabaseDriver = "postgres"
	MySQL    DatabaseDriver = "mysql"
	Static   DatabaseDriver = "static"
)

type DatabaseConnection struct {
	Driver   DatabaseDriver
	Postgres *gorm.DB
	MySQL    *gorm.DB
}

func NewDatabaseConnection(configs *configs.AppConfig) *DatabaseConnection {
	var db DatabaseConnection

	switch configs.Database.DRIVER {
	case "postgres":
		db.Driver = Postgres
		db.Postgres = newPostgres(configs)
	case "mysql":
		db.Driver = MySQL
		db.MySQL = newMySQL(configs)
	case "static":
		db.Driver = Static
	}

	return &db
}

func newPostgres(configs *configs.AppConfig) *gorm.DB {
	var connectionString string

	switch configs.Database.DRIVER {
	case "postgres":
		connectionString = fmt.Sprintf("%s://%s:%s@%s:%s/%s",
			configs.Database.DRIVER,
			configs.Database.USERNAME,
			configs.Database.PASSWORD,
			configs.Database.HOST,
			configs.Database.PORT,
			configs.Database.DATABASE)
	}

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}

func newMySQL(configs *configs.AppConfig) *gorm.DB {
	var connectionString string

	switch configs.Database.DRIVER {
	case "mysql":
		connectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
			configs.Database.USERNAME,
			configs.Database.PASSWORD,
			configs.Database.HOST,
			configs.Database.PORT,
			configs.Database.DATABASE)
	}

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}

func (db *DatabaseConnection) CloseConnection() {
	if db.Postgres != nil {
		db, _ := db.Postgres.DB()
		db.Close()
	}
	if db.MySQL != nil {
		db, _ := db.Postgres.DB()
		db.Close()
	}
}
