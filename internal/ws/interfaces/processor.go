package interfaces

import (
	"messenger/internal/messaging/interfaces"

	"github.com/gorilla/websocket"
)

type WebSocketProcessor interface {
	interfaces.MessageProcessor
	SetConnection(connection *websocket.Conn)
}
