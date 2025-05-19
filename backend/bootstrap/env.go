package bootstrap

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Env struct {
	AppEnv        string `mapstructure:"APP_ENV"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	DBHost        string `mapstructure:"DB_HOST"`
	DBPort        string `mapstructure:"DB_PORT"`
	DBUser        string `mapstructure:"DB_USER"`
	DBPass        string `mapstructure:"DB_PASS"`
	DBName        string `mapstructure:"DB_NAME"`
}

func NewEnv() *Env {
	env := Env{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Println("No .env file found")

		viper.AutomaticEnv()

		err = viper.Unmarshal(&env)
		if err != nil {
			log.Fatal("Environment can't be loaded: ", err)
		}
		return &env
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	if env.AppEnv == "development" {
		log.Info("The App is running in development env")
	}

	return &env
}
