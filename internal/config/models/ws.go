package models

import (
	"fmt"
	"net"
	"strconv"
)

type WebSocket struct {
	Host           string   `mapstructure:"host"`
	Port           string   `mapstructure:"port"`
	Debug          bool     `mapstructure:"debug"`
	InvalidOrigins []string `mapstructure:"invalid_origins"`
}

// Validate проверяет конфигурацию WebSocket на корректность.
// Что:
// - Поле Host не пустое и содержит валидный IP-адрес.
// - Поле Port не пустое, является числом и находится в диапазоне от 1 до 65535.
// - Поле InvalidOrigins не пустое.
// Если какое-либо из этих условий не выполнено, возвращается ошибка.
func (ws *WebSocket) Validate() error {
	if ws.Host == "" {
		return fmt.Errorf("host обязателен")
	}
	if net.ParseIP(ws.Host) == nil {
		return fmt.Errorf("host должен быть валидным IP-адресом")
	}
	if ws.Port == "" {
		return fmt.Errorf("port обязателен")
	}
	port, err := strconv.Atoi(ws.Port)
	if err != nil || port <= 0 || port > 65535 {
		return fmt.Errorf("port должен быть числом в диапазоне от 1 до 65535")
	}
	if len(ws.InvalidOrigins) == 0 {
		return fmt.Errorf("invalid_origins не может быть пустым")
	}
	return nil
}
