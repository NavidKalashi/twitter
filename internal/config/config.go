package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Twitter string `mapstructure:"twitter"`
	Port    int    `mapstructure:"port"`
	DB      struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Name     string `mapstructure:"name"`
	} `mapstructure:"db"`
}

var Cfg Config

func LoadConfig() (*Config, error){
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("internal/config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error reading config file, %s", err)
	}

	
	if err := viper.Unmarshal(&Cfg); err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	return &Cfg, nil
}