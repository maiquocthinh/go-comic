package config

import (
	"errors"
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server  ServerConfig
	MySQL   MySQLConfig
	Redis   RedisConfig
	Dropbox DropboxConfig
}

type ServerConfig struct {
	Port              string
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	JwtSecretKey      string
	JwtUserExpiration int // second
}

type MySQLConfig struct {
	Host       string
	Port       int
	Username   string
	Password   string
	DBName     string
	DriverName string
}

type RedisConfig struct {
	Addr        string
	Username    string
	Password    string
	DB          int
	PoolSize    int
	PoolTimeout int
}

type DropboxConfig struct {
	AppKey       string
	AppSecret    string
	RefreshToken string
}

func NewConfig(filename string) (*Config, error) {
	v := viper.New()

	// load config
	v.SetConfigName(filename)
	v.SetConfigType("yml")
	v.AddConfigPath(".")
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found!")
		}
		return nil, err
	}

	// parse config
	var config Config
	if err := v.Unmarshal(&config); err != nil {
		log.Printf("unable to decode into struct!")
		return nil, err
	}

	return &config, nil
}
