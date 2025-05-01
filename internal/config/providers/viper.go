package providers

import (
	"messager/internal/config/models"

	"github.com/spf13/viper"
)

type ViperConfigProvider struct{}

// Load загружает файл конфигурации с использованием библиотеки Viper и преобразует его содержимое
// в структуру Config. Также выполняется проверка загруженной конфигурации.
//
// Параметры:
//   - path: Путь к директории, где находится файл конфигурации.
//   - filename: Имя файла конфигурации (без расширения).
//   - configType: Тип файла конфигурации (например, "json", "yaml").
//
// Возвращает:
//   - *models.Config: Указатель на загруженную и проверенную структуру конфигурации.
//   - error: Ошибка, если файл конфигурации не может быть прочитан, преобразован или проверен.
func (v *ViperConfigProvider) Load(path, filename, configType string) (*models.Config, error) {
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

	if err := config.Validate(); err != nil {
		return nil, err
	}

	return config, nil
}
