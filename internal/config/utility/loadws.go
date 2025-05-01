package utility

import (
	conf "messager/internal/config/models"
)

// LoadWebsocketConfig загружает конфигурацию WebSocket из предоставленного объекта webSocketConfig.
// Она извлекает хост, порт, флаг debug и список недопустимых источников для WebSocket.
//
// Параметры:
//   - webSocketConfig: Указатель на объект conf.WebSocket, содержащий конфигурацию WebSocket.
//
// Возвращает:
//   - string: Хост WebSocket.
//   - string: Порт WebSocket.
//   - bool: Флаг debug WebSocket.
//   - []string: Список недопустимых источников WebSocket.
func LoadWebsocket(webSocketConfig conf.WebSocket) (string, string, bool, []string) {
	wsHost := webSocketConfig.Host
	wsPort := webSocketConfig.Port
	wsDebug := webSocketConfig.Debug
	invalidOrigins := webSocketConfig.InvalidOrigins

	return wsHost, wsPort, wsDebug, invalidOrigins
}
