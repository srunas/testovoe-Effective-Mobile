package config

import (
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

type ServerConfig struct {
	Port int `mapstructure:"port"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
	SSLMode  string `mapstructure:"ssl_mode"`
}

//nolint:gochecknoglobals // singleton паттерн для конфига
var (
	once     sync.Once
	instance *Config
)

func Get() *Config {
	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("internal/config")
		viper.AddConfigPath(".")

		if err := viper.ReadInConfig(); err != nil {
			panic("ошибка чтения конфига: " + err.Error())
		}

		instance = &Config{}
		if err := viper.Unmarshal(instance); err != nil {
			panic("ошибка парсинга конфига: " + err.Error())
		}
	})
	return instance
}
