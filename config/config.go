package config

import (
	"github.com/spf13/viper"
)

// Config holds the structure of the config file
type Config struct {
	Database struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Name     string `mapstructure:"name"`
	}
}

// LoadConfig loads the config.yaml file and unmarshals it into the Config struct
func LoadConfig() (*Config, error) {
	var config Config

	// Set the config file name and path
	viper.SetConfigName("config")   // name of config file (without extension)
	viper.AddConfigPath("./config") // path to look for the config file in

	// Read in the config file
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	// Unmarshal the config file into the Config struct
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
