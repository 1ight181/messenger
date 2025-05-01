package interfaces

import "messager/internal/config/models"

type ConfigProvider interface {
	Load(path, filename, configType string) (*models.Config, error)
}
