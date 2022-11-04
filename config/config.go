package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	ADDRESS                   string        `mapstructure:"USOS_HTTP_ADDRESS"`

	DB_HOST                string        `mapstructure:"USOS_DB_HOST"`
	DB_PORT                string        `mapstructure:"USOS_DB_PORT"`
	DB_USER                string        `mapstructure:"USOS_DB_USER"`
	DB_PASSWORD            string        `mapstructure:"USOS_DB_PASSWORD"`
	DB_DBNAME              string        `mapstructure:"USOS_DB_DBNAME"`
}

func LoadConfigFromFile(configFileName string) (config Config, err error) {
	viper.SetConfigFile(configFileName)

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
