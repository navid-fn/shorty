package utils

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBHost     string `mapstructure:"DBHOST"`
	DBname     string `mapstructure:"DBNAME"`
	DBpassword string `mapstructure:"DBPASSWORD"`
	DBport     string `mapstructure:"DBPORT"`
	DBusername string `mapstructure:"DBUSERNAME"`
}

func LoadConfig() (c Config, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return
		}
	}

	err = viper.Unmarshal(&c)
	return
}
