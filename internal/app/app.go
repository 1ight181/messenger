package app

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"

	confloader "messenger/internal/config/loaders"
	confrpov "messenger/internal/config/providers"
	"messenger/internal/factories"

	processor "messenger/internal/messaging/processor"
	"messenger/internal/messaging/receiver"
	"messenger/internal/messaging/sender"

	ws "messenger/internal/ws"
	wsupgr "messenger/internal/ws/upgraders"

	"github.com/spf13/viper"
)

// Run инициализирует и запускает приложение WebSocket-сервера.
// Выполняются следующие шаги:
// 1. Загружается конфигурационный файл с использованием ViperConfigProvider.
//   - Если переменная окружения CONFIG_PATH не задана, используется путь по умолчанию.
//
// 2. Обрабатываются возможные ошибки при загрузке конфигурации, включая:
//   - Ошибки пути
//   - Отсутствие конфигурационного файла
//   - Ошибки разбора конфигурации
//
// 3. Загружается TLS-сертификат для безопасной связи.
// 4. Настраиваются параметры WebSocket, включая хост, порт, режим отладки и недопустимые источники.
// 5. Создается и инициализируется обработчик WebSocket с необходимыми компонентами:
//   - Upgrader для WebSocket-соединений
//   - Отправитель, получатель и обработчик сообщений.
//
// 6. Настраивается HTTP-сервер с конфигурацией TLS и обработчиком WebSocket.
// 7. Запускается WebSocket-сервер.
//
// Эта функция регистрирует фатальные ошибки и завершает приложение, если возникают
// критические проблемы во время инициализации или запуска.
func Run() {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "../configs"
		log.Printf("CONFIG_PATH не задан, используется путь по умолчанию: %s", configPath)
	}

	configFilename := "config"
	configFiletype := "yaml"

	viperConfigProvider := &confrpov.ViperConfigProvider{}

	config, err := confloader.LoadConfig(
		viperConfigProvider,
		configPath,
		configFilename,
		configFiletype,
	)

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

	certificate, err := confloader.LoadCertificate(config.Certificate)
	if err != nil {
		log.Fatalf("Ошибка загрузки сертификата: %v", err)
	}

	wsHost, wsPort, wsDebug, invalidOrigins := confloader.LoadWebsocket(config.WebSocket)

	wsUpgrager := wsupgr.NewUpgrader(wsDebug, invalidOrigins)

	wsProcessorOptions :=
		processor.Options{
			ErrorResponseText: "Ошибка получена и обработана",
			InfoResponseText:  "Информационное собщение получено и обработано",
			DataResponseText:  "Сообщение с данными получено и обработано",
		}

	wsSenderOptions := sender.Options{}
	wsReceiverOptions := receiver.Options{}

	webSocketHandlerFactory := factories.New(
		wsUpgrager,
		wsSenderOptions,
		wsReceiverOptions,
		wsProcessorOptions,
	)

	wsHandlerFunc := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler := webSocketHandlerFactory.NewHandler()
		handler.HandleWebSocket(w, r)
	})

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
