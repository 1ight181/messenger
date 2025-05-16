package app

import (
	"log"

	viperprov "messenger/internal/config/providers/viper"

	processor "messenger/internal/messaging/processor"
	"messenger/internal/messaging/receiver"
	"messenger/internal/messaging/sender"
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
	appConfigOptions := AppConfigOptions{
		Provider:    &viperprov.ViperConfigProvider{},
		FileName:    "config",
		FileType:    "yaml",
		EnvVar:      "CONFIG_PATH",
		DefaultPath: "../config",
	}

	config, err := loadAppConfig(appConfigOptions)
	if err != nil {
		log.Fatalf("Ошибка загрузки конфигурации: %v", err)
	}

	tlsConfig, err := loadAppCertificateConfig(config.Certificate)
	if err != nil {
		log.Fatalf("Ошибка загрузки сертификата: %v", err)
	}

	wsProcessorOptions :=
		processor.Options{
			ErrorResponseText: "Ошибка получена и обработана",
			InfoResponseText:  "Информационное собщение получено и обработано",
			DataResponseText:  "Сообщение с данными получено и обработано",
		}

	wsSenderOptions := sender.Options{}
	wsReceiverOptions := receiver.Options{}

	webSocketServiceOptions := WebSocketServiceOptions{
		Config:           config.WebSocket,
		TLSConfig:        &tlsConfig,
		SenderOptions:    wsSenderOptions,
		ReceiverOptions:  wsReceiverOptions,
		ProcessorOptions: wsProcessorOptions,
	}
	wsService := loadAppWebSocketService(webSocketServiceOptions)

	wsService.StartServer()
}
