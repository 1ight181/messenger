package loaders

import (
	"messenger/internal/config/models"
	"messenger/internal/config/providers/interfaces"
)

// LoadConfig загружает конфигурацию с использованием предоставленного ConfigProvider.
// Он принимает следующие параметры:
// - configProvider: Реализация интерфейса ConfigProvider, используемая для загрузки конфигурации.
// - path: Директория, где находится файл конфигурации.
// - filename: Имя файла конфигурации.
// - configType: Тип файла конфигурации (например, "json", "yaml").
// Возвращает указатель на модель Config и ошибку, если процесс загрузки завершился неудачей.
func LoadConfig(configProvider interfaces.ConfigProvider, path, filename, configType string) (*models.Config, error) {
	return configProvider.Load(path, filename, configType)
}
