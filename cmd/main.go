package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"

	confutil "messager/internal/config/loaders"
	confrpov "messager/internal/config/providers"

	processor "messager/internal/messaging/processor"
	receiver "messager/internal/messaging/receiver"
	sender "messager/internal/messaging/sender"

	ws "messager/internal/ws"
	wsh "messager/internal/ws/handlers"
	wsutil "messager/internal/ws/upgraders"

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
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "../configs"
		log.Printf("CONFIG_PATH не задан, используется путь по умолчанию: %s", configPath)
	}

	configFilename := "config"
	configFiletype := "yaml"

	viperConfigProvider := &confrpov.ViperConfigProvider{}

	config, err := confutil.LoadConfig(viperConfigProvider, configPath, configFilename, configFiletype)

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

	certificate, err := confutil.LoadCertificate(config.Certificate)
	if err != nil {
		log.Fatalf("Ошибка загрузки сертификата: %v", err)
	}

	wsHost, wsPort, wsDebug, invalidOrigins := confutil.LoadWebsocket(config.WebSocket)

	wsUpgrager := wsutil.NewUpgrader(wsDebug, invalidOrigins)

	wsHandler := wsh.NewWebSocketHandler(
		wsUpgrager,
		sender.NewWebSocketMessageSender(),
		receiver.NewWebSocketMessageReceiver(),
		processor.NewWebSocketMessageProcessor(),
	)

	wsHandlerFunc := http.HandlerFunc(wsHandler.HandleWebSocket)

	address := fmt.Sprintf("%s:%s", wsHost, wsPort)

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{certificate},
		MinVersion:   tls.VersionTLS12,
	}

	httpServer := &http.Server{
		Addr:      address,
		Handler:   wsHandlerFunc,
		TLSConfig: tlsConfig,
	}

	wsService := ws.NewWebsocketService(
		httpServer,
	)

	wsService.StartServer()
}
