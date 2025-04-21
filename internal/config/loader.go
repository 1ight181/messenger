package config_reader

import (
	models "messager/internal/config/models"

	"github.com/spf13/viper"
)

func LoadConfig(path, filename, configType string) (*models.Config, error) {
	viper.SetConfigName(filename)
	viper.AddConfigPath(path)
	viper.SetConfigType(configType)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	config := &models.Config{}
	if err := viper.Unmarshal(config); err != nil {
		return nil, err
	}

	return config, nil
}
