package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type ViperConfigProvider struct {
	config Config
}

// C# style config loading with APP_ENV environment variable
func NewViperConfigProvider(configPath string) (ConfigProvider, error) {
	env := os.Getenv("APP_ENV")

	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(configPath)

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	// Override with environment-specific config
	viper.SetConfigName(fmt.Sprintf("config.%s", env))
	if err := viper.MergeInConfig(); err != nil {
		fmt.Printf("No environment-specific config found for '%s': %v\n", env, err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unable to decode into struct: %w", err)
	}

	return &ViperConfigProvider{config: cfg}, nil
}

func (v *ViperConfigProvider) GetConfig() Config {
	return v.config
}
