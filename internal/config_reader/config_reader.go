package config_reader

import (
	configTypes "messager/internal/config_reader/config_types"

	"github.com/spf13/viper"
)

func LoadConfig(path, filename, configType string) (*configTypes.Config, error) {
	viper.SetConfigName(filename)
	viper.AddConfigPath(path)
	viper.SetConfigType(configType)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	config := &configTypes.Config{}
	if err := viper.Unmarshal(config); err != nil {
		return nil, err
	}

	return config, nil
}
