package interfaces

import (
	"messenger/internal/messaging/interfaces"

	"github.com/gorilla/websocket"
)

type WebSocketSender interface {
	interfaces.MessageSender
	SetConnection(connection *websocket.Conn)
}
