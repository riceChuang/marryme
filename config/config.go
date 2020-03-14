package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	configuration *Config
)

func Initial() *Config {
	viper.SetConfigName(".env")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Panicf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Panicf("Unable to decode config into struct, %v", err)
	}

	return configuration
}

func GetConfig() *Config {
	return configuration
}
