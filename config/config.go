package config

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Val Val
var Val Config

// Config Config
type Config struct {
	Mode string `mapstructure:"MODE"`

	Domain string `mapstructure:"DOMAIN"`

	Port string `mapstructure:"PORT"`

	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_User"`
	DBPass     string `mapstructure:"DB_Pass"`
	DBDatabase string `mapstructure:"DB_DATABASE"`

	DBMaxConn     int `mapstructure:"DB_MAX_CONN"`
	DBIdleConn    int `mapstructure:"DB_IDLE_CONN"`
	DBMaxLifeTime int `mapstructure:"DB_MAX_LIFE_TIME"`

	RedisHost string `mapstructure:"REDIS_HOST"`
	RedisPort string `mapstructure:"REDIS_PORT"`
	RedisUser string `mapstructure:"REDIS_USER"`
	RedisPass string `mapstructure:"REDIS_PASS"`

	GoogleSecretKey string `mapstructure:"GOOGLE_SECRET_KEY"`
	GoogleClientID  string `mapstructure:"GOOLE_CLIENT_ID"`
	RedirectURL     string `mapstructure:"REDIRECT_URL"`

	TimeFormat string `mapstructure:"TIMEFORMAT"`

	JWTTokenLife int    `mapstructure:"JWT_TOKEN_LIFE"`
	JWTSecret    string `mapstructure:"JWT_SECRET"`
}

// Init Init
func Init() {
	// 讀config.yaml
	viper.SetConfigName("config.yaml")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %v ", err))
	}
	if err := viper.Unmarshal(&Val); err != nil {
		panic(fmt.Errorf("unable to decode into struct, %v", err))
	}

	log.WithFields(log.Fields{
		"val": Val,
	}).Info("config loaded")
}
