package config

import (
	models "messager/internal/config/models"

	"github.com/spf13/viper"
)

// LoadConfig загружает конфигурацию из указанного файла и преобразует её в структуру Config.
//
// Параметры:
//   - path: Путь к директории, где находится файл конфигурации.
//   - filename: Имя файла конфигурации (без расширения).
//   - configType: Тип файла конфигурации (например, "json", "yaml").
//
// Возвращает:
//   - *models.Config: Указатель на структуру Config, заполненную данными конфигурации.
//   - error: Ошибка, если файл конфигурации не удалось прочитать или преобразовать.
//
// Эта функция использует библиотеку Viper для чтения и разбора файла конфигурации.
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
