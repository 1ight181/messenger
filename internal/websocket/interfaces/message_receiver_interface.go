package interfaces

import (
	"messager/internal/messaging/receiver"

	"github.com/gorilla/websocket"
)

type WebSocketReceiver interface {
	receiver.MessageReceiver
	SetConnection(connection *websocket.Conn)
}
