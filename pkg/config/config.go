package config

import (
	"errors"

	"github.com/spf13/viper"
)

// Config
type Config struct {
	ServerConfig ServerConfig `yaml:"ServerConfig"`
}

// ServerConfig
type ServerConfig struct {
	Port                int    `yaml:"Port"`
	TimeoutSecs         int    `yaml:"TimeoutSecs"`
	ReadTimeoutSecs     int    `yaml:"ReadTimeoutSecs"`
	WriteTimeoutSecs    int    `yaml:"WriteTimeoutSecs"`
	AppVersion          string `yaml:"AppVersion"`
	Mode                string `yaml:"Mode"`
	RoutePrefix         string `yaml:"RoutePrefix"`
	Debug               bool   `yaml:"Debug"`
	ShutdownTimeoutSecs int    `yaml:"ShutdownTimeoutSecs"`
}

// LoadConfig reads configuration from a file
func LoadConfig(fileName string) (*Config, error) {
	v := viper.New()

	v.SetConfigName(fileName)
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}
	var c Config
	err := v.Unmarshal(&c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
