package bootstrap

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Env struct {
	AppEnv        string `mapstructure:"APP_ENV"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	ServerPort    string `mapstructure:"SERVER_PORT"`
	DBHost        string `mapstructure:"DB_HOST"`
	DBPort        string `mapstructure:"DB_PORT"`
	DBUser        string `mapstructure:"DB_USER"`
	DBPass        string `mapstructure:"DB_PASS"`
	DBName        string `mapstructure:"DB_NAME"`
	InfuraKey 	  string `mapstructure:"INFURA_KEY"`
}

func NewEnv() *Env {
	env := Env{}

	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Warn(".env file not found, using system environment variables")
	}

	if err := viper.Unmarshal(&env); err != nil {
		log.Fatal("Environment variables can't be loaded: ", err)
	}

	if env.AppEnv == "development" {
		log.Info("The App is running in development environment")
	}

	return &env
}
