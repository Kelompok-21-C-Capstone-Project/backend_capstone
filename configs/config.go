package configs

import (
	"sync"

	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

type AppConfig struct {
	App struct {
		NAME string `mapstructure:"APP_NAME"`
		PORT int    `mapstructure:"APP_PORT"`
		JWT  string `mapstructure:"JWT_KEY"`
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

	defaultConfig.App.PORT = 5404

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
