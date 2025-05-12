package interfaces

import (
	"messenger/internal/messaging/interfaces"

	"github.com/gorilla/websocket"
)

type WebSocketReceiver interface {
	interfaces.MessageReceiver
	SetConnection(connection *websocket.Conn)
}
