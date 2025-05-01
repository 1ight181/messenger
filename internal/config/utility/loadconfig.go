package utility

import (
	"messager/internal/config/models"
	"messager/internal/config/providers/interfaces"
)

func LoadConfig(configProvider interfaces.ConfigProvider, path, filename, configType string) (*models.Config, error) {
	return configProvider.Load(path, filename, configType)
}
