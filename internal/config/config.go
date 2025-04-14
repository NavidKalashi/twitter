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
	Redis struct {
		Addr     string `mapstructure:"addr"`
		Password string `mapstructure:"password"`
		DB       int    `mapstructure:"db"`
	} `mapstructure:"redis"`
	Minio struct {
		Endpoint        string `mapstructure:"endpoint"`
		AccessKeyID     string `mapstructure:"accessKeyID"`
		SecretAccessKey string `mapstructure:"secretAccessKey"`
		UseSSL          bool   `mapstructure:"useSSL"`
		BucketName      string `mapstructure:"bucketName"`
	} `mapstructure:"minio"`
	RabbitMQ struct {
		URL string `mapstructure:"url"`
	} `mapstructure:"rabbitmq`
}

var Cfg Config

func LoadConfig() (*Config, error) {
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
