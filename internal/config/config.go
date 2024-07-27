package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	//TODO: create a config struct
}

func LoadConfig() *Config {

	var cfg Config

	myViper := viper.New()

	myViper.SetConfigName("config")
	myViper.SetConfigType("yaml")

	myViper.AddConfigPath(".")

	if err := myViper.ReadInConfig(); err != nil {
		log.Fatal("Failed reading config file")
	}

	if err := myViper.Unmarshal(&cfg); err != nil {
		log.Fatal("The configuration file in struct could not be parsed")
	}

	return &cfg
}
