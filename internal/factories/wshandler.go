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

func (f *WebSocketHandlerFactory) NewHandler() *handlers.WebSocketHandler {
	return handlers.New(
		f.upgrader,
		sender.New(f.senderOptions),
		receiver.New(f.receiverOptions),
		processor.New(f.processorOptions),
	)
}
