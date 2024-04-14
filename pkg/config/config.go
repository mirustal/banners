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
	TestDBName string `mapstructure:"POSTGRES_TEST_DB"`
	Ip			   	string `mapstructure:"SERVICE_IP" default=0.0.0.0`
	Port			string `mapstructure:"SERVICE_PORT" default=8000`

	JwtSecret    string        `mapstructure:"JWT_SECRET"`
	JwtExpiresIn time.Duration `mapstructure:"JWT_EXPIRED_IN"`
	JwtMaxAge    int           `mapstructure:"JWT_MAXAGE"`

	ClientOrigin string `mapstructure:"CLIENT_ORIGIN"`
	LogDB bool `mapstructure: LOG_DB`
}



func GetConfig() (config *Config) {
	viper.AddConfigPath(".")
	viper.SetConfigType("env")
	viper.SetConfigName("app")


	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("config path not readssss %s ", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil{
		log.Fatal("config bad unmarshal %f", err)
	}
	return config
}

