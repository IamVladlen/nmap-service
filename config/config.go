package config

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/spf13/viper"
)

// Config stores data from /config/config.yaml and /.env files.
type Config struct {
	App struct {
		LogLevel string `mapstructure:"log_level"`
	}

	GRPC struct {
		PORT string `mapstructure:"GRPC_NMAP_PORT"`
	} `mapstructure:",squash"`
}

// New creates new config instance with data from
// config files.
func New() (*Config, error) {
	var cfg *Config

	viper.SetConfigFile("config/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		return &Config{}, fmt.Errorf("config - New: %w", err)
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return &Config{}, fmt.Errorf("config - New: %w", err)
	}

	loadEnv(cfg)

	return cfg, nil
}

func loadEnv(cfg *Config) {
	cfg.GRPC.PORT = os.Getenv("GRPC_NMAP_PORT")
}
