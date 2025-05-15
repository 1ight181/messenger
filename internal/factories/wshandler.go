package factories

import (
	"messenger/internal/messaging/processor"
	"messenger/internal/messaging/receiver"
	"messenger/internal/messaging/sender"

	"messenger/internal/ws/handlers"

	"github.com/gorilla/websocket"
)

type WebSocketHandlerFactory struct {
	upgrader         websocket.Upgrader
	senderOptions    sender.Options
	receiverOptions  receiver.Options
	processorOptions processor.Options
}

// New создает и возвращает новый экземпляр WebSocketHandlerFactory с предоставленными
// параметрами upgrader, senderOptions, receiverOptions и processorOpts.
// Инициализирует фабрику необходимыми зависимостями для обработки WebSocket-соединений.
//
// Параметры:
//
//   - upgrader        - websocket.Upgrader для апгрейда HTTP-соединений до WebSocket.
//   - senderOptions   - Опции конфигурации для компонента отправки сообщений.
//   - receiverOptions - Опции конфигурации для компонента приема сообщений.
//   - processorOpts   - Опции конфигурации для компонента обработки сообщений.
//
// Возвращает:
//
//   - *WebSocketHandlerFactory - Указатель на инициализированную фабрику.
func New(
	upgrader websocket.Upgrader,
	senderOptions sender.Options,
	receiverOptions receiver.Options,
	processorOpts processor.Options,
) *WebSocketHandlerFactory {
	return &WebSocketHandlerFactory{
		upgrader:         upgrader,
		senderOptions:    senderOptions,
		receiverOptions:  receiverOptions,
		processorOptions: processorOpts,
	}
}

// NewHandler создает и возвращает новый экземпляр handlers.WebSocketHandler,
// инициализируя его настроенным upgrader, sender, receiver и processor.
// Зависимости создаются с использованием опций фабрики.
func (f *WebSocketHandlerFactory) NewHandler() *handlers.WebSocketHandler {
	return handlers.New(
		f.upgrader,
		sender.New(f.senderOptions),
		receiver.New(f.receiverOptions),
		processor.New(f.processorOptions),
	)
}
