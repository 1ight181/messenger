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
