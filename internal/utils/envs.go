package utils

import (
	"fmt"

	"github.com/spf13/viper"
)


func LoadDBConfig() (c string, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return "", err
	}

	dbUser := viper.Get("DBUSERNAME").(string)
	dbPassword := viper.Get("DBPASSWORD").(string)
	dbName := viper.Get("DBNAME").(string)
	dbPort := viper.Get("DBPORT").(string)
	dbHost := viper.Get("DBHOST").(string)

	dbConfig := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName, dbPort)
	return dbConfig, nil
}
