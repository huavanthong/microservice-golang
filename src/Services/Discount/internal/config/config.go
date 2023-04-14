package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type AppConfig struct {
	Port string `mapstructure:"port"`
}
type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
}

type LoggerConfig struct {
	LogLevel string `mapstructure:"log_level"`
}
type Config struct {
	App      AppConfig      `mapstructure:"App"`
	Database DatabaseConfig `mapstructure:"database"`
	Logger   LoggerConfig   `mapstructure:"logger"`
}

func LoadConfig(path string) (*Config, error) {
	// set default values for config
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "5432")
	viper.SetDefault("database.user", "postgres")
	viper.SetDefault("database.password", "")
	viper.SetDefault("database.dbname", "discount_service")

	viper.SetDefault("logger.loglevel", "info")

	// set config file name and path
	//viper.AddConfigPath(path)
	viper.SetConfigFile(path)

	// read config file
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("Config file not found; using default values.")
		} else {
			return nil, err
		}
	}

	// unmarshal config into Config struct
	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	// Check error
	if config.Database.Host == "" {
		return nil, fmt.Errorf("missing 'host' field in database config")
	}
	if config.Database.Port == "" {
		return nil, fmt.Errorf("missing 'port' field in database config")
	}
	if config.Database.User == "" {
		return nil, fmt.Errorf("missing 'user' field in database config")
	}
	if config.Database.Password == "" {
		return nil, fmt.Errorf("missing 'password' field in database config")
	}
	if config.Database.DBName == "" {
		return nil, fmt.Errorf("missing 'dbname' field in database config")
	}

	return &config, nil
}