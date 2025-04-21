package main

import (
	"log"
	"os"

	config "messager/internal/config"
	"messager/internal/messaging/processor"
	"messager/internal/messaging/receiver"
	"messager/internal/messaging/sender"
	"messager/internal/websocket"

	"github.com/spf13/viper"
)

// main — это точка входа в приложение. Она инициализирует конфигурацию,
// настраивает WebSocket-сервис и запускает WebSocket-сервер.
//
// Функция выполняет следующие шаги:
//  1. Загружает конфигурационный файл из указанного пути, с заданным именем и типом файла.
//  2. Обрабатывает возможные ошибки при загрузке конфигурации, включая ошибки пути,
//     отсутствие конфигурационного файла и ошибки разбора.
//  3. Извлекает значения конфигурации, связанные с WebSocket, такие как хост, порт и режим отладки.
//  4. Инициализирует компоненты для отправки, получения и обработки WebSocket-сообщений.
//  5. Создает WebSocket-сервис, используя инициализированные компоненты и значения конфигурации.
//  6. Запускает WebSocket-сервер на указанном хосте и порту.
func main() {
	path := "../configs"
	configFilename := "config"
	configFiletype := "yaml"

	config, err := config.LoadConfig(path, configFilename, configFiletype)

	if err != nil {
		switch e := err.(type) {
		case *os.PathError:
			log.Fatalf("Ошибка пути: %v", e)
		case *viper.ConfigFileNotFoundError:
			log.Fatalf("Конфигурационный файл не найден: %v", e)
		case *viper.ConfigParseError:
			log.Fatalf("Ошибка при разборе конфигурации: %v", e)
		default:
			log.Fatalf("Неизвестная ошибка: %v", e)
		}
	}

	wsHost := config.WebSocket.Host
	wsPort := config.WebSocket.Port
	wsDebug := config.WebSocket.Debug

	messageSender := sender.NewWebSocketMessageSender()
	messageReceiver := receiver.NewWebSocketMessageReceiver()
	messageProcessor := processor.NewWebSocketMessageProcessor()

	wsService := websocket.NewWebsocketService(
		wsDebug,
		wsHost,
		messageSender,
		messageReceiver,
		messageProcessor,
	)

	wsService.StartServer(wsHost, wsPort)
}
