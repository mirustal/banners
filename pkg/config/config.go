package config

import (
	"log"
	"time"
	"github.com/spf13/viper"
)

type Config struct {
	DBHost         	string `mapstructure:"POSTGRES_HOST" env:"DBHost`
	DBUserName     	string `mapstructure:"POSTGRES_USER"`
	DBUserPassword 	string `mapstructure:"POSTGRES_PASSWORD"`
	DBName         	string `mapstructure:"POSTGRES_DB"`
	DBPort         	string `mapstructure:"POSTGRES_PORT"`
	Ip			   	string `mapstructure:"SERVICE_IP" default=0.0.0.0`
	Port			string `mapstructure:"SERVICE_PORT" default=8000`

	JwtSecret    string        `mapstructure:"JWT_SECRET"`
	JwtExpiresIn time.Duration `mapstructure:"JWT_EXPIRED_IN"`
	JwtMaxAge    int           `mapstructure:"JWT_MAXAGE"`

	ClientOrigin string `mapstructure:"CLIENT_ORIGIN"`
}



func GetConfig() (config *Config) {
	path := "/Users/mirustal/Documents/project/go/avito_tech/"
	if path == "" {
		log.Fatal("config path is empty")
	}

	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("config path not read")
	}

	err = viper.Unmarshal(&config)
	return config
}

