package configs

import (
	"os"
	"strconv"
	"sync"

	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

type AppConfig struct {
	App struct {
		NAME      string `mapstructure:"APP_NAME"`
		PORT      int    `mapstructure:"APP_PORT"`
		JWT       string `mapstructure:"JWT_KEY"`
		ADMIN_KEY string `mapstructure:"APP_ADMINKEY"`
	}
	Database struct {
		DRIVER     string `mapstructure:"DB_DRIVER"`
		CONNECTION string `mapstructure:"DB_CONNECTION"`
		HOST       string `mapstructure:"DB_HOST"`
		PORT       string `mapstructure:"DB_PORT"`
		DATABASE   string `mapstructure:"DB_DATABASE"`
		USERNAME   string `mapstructure:"DB_USERNAME"`
		PASSWORD   string `mapstructure:"DB_PASSWORD"`
	}
	API_Midtrans struct {
		SERVER_KEY string `mapstructure:"MIDTRANS_SERVER_KEY"`
		ENV        string `mapstructure:"MIDTRANS_ENV"`
	}
	API_Mailjet struct {
		PRIVATE_KEY string `mapstructure:"MAILJET_PRIVATE_KEY"`
		PUBLIC_KEY  string `mapstructure:"MAILJET_PUBLIC_KEY"`
	}
}

var lock = &sync.Mutex{}
var appConfig *AppConfig

func GetConfig() *AppConfig {
	lock.Lock()
	defer lock.Unlock()

	if appConfig == nil {
		appConfig = initConfig()
	}

	return appConfig
}

func initConfig() *AppConfig {
	var defaultConfig AppConfig

	defaultConfig.App.PORT, _ = strconv.Atoi(os.Getenv("PORT"))
	defaultConfig.App.NAME = os.Getenv("APP_NAME")
	defaultConfig.App.JWT = os.Getenv("JWT")
	defaultConfig.App.ADMIN_KEY = os.Getenv("ADMIN_KEY")

	defaultConfig.Database.DRIVER = os.Getenv("DB_DRIVER")
	defaultConfig.Database.CONNECTION = os.Getenv("DATABASE_URL")
	defaultConfig.Database.HOST = os.Getenv("DB_HOST")
	defaultConfig.Database.PORT = os.Getenv("DB_PORT")
	defaultConfig.Database.DATABASE = os.Getenv("DB_DATABASE")
	defaultConfig.Database.USERNAME = os.Getenv("DB_USERNAME")
	defaultConfig.Database.PASSWORD = os.Getenv("DB_PASSWORD")

	defaultConfig.API_Midtrans.SERVER_KEY = os.Getenv("MIDTRANS_KEY")
	defaultConfig.API_Midtrans.ENV = os.Getenv("MIDTRANS_ENV")

	defaultConfig.API_Mailjet.PRIVATE_KEY = os.Getenv("MAILJET_PRIVATE_KEY")
	defaultConfig.API_Mailjet.PUBLIC_KEY = os.Getenv("MAILJET_PUBLIC_KEY")

	viper.SetConfigType("env")
	viper.SetConfigName("config")
	viper.AddConfigPath("./configs/")

	if err := viper.ReadInConfig(); err != nil {
		log.Info("error when load config file", err)
		return &defaultConfig
	}

	var finalConfig AppConfig

	errApp := viper.Unmarshal(&finalConfig.App)
	errDB := viper.Unmarshal(&finalConfig.Database)
	errMidtrans := viper.Unmarshal(&finalConfig.API_Midtrans)

	if errApp != nil || errDB != nil {
		if errApp != nil {
			log.Info("error when parse config file", errApp)
		}
		if errDB != nil {
			log.Info("error when parse config file", errDB)
		}
		if errMidtrans != nil {
			log.Info("error when parse config file", errDB)
		}
		return &defaultConfig
	}

	return &finalConfig
}
