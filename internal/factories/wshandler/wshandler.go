package websocket

import (
	"messenger/internal/messaging/processor"
	"messenger/internal/messaging/receiver"
	"messenger/internal/messaging/sender"

	"messenger/internal/ws/handlers"

	"github.com/gorilla/websocket"
)

type WebSocketHandlerFactory struct {
	options Options
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
	options Options,
) *WebSocketHandlerFactory {
	return &WebSocketHandlerFactory{
		options: options,
	}
}

type Options struct {
	Upgrader         websocket.Upgrader
	SenderOptions    sender.Options
	ReceiverOptions  receiver.Options
	ProcessorOptions processor.Options
}

// NewHandler создает и возвращает новый экземпляр handlers.WebSocketHandler,
// инициализируя его настроенным upgrader, sender, receiver и processor.
// Зависимости создаются с использованием опций фабрики.
func (f *WebSocketHandlerFactory) NewHandler() *handlers.WebSocketHandler {
	return handlers.New(
		f.options.Upgrader,
		sender.New(f.options.SenderOptions),
		receiver.New(f.options.ReceiverOptions),
		processor.New(f.options.ProcessorOptions),
	)
}
