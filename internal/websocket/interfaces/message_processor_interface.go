package interfaces

import (
	"messager/internal/messaging/processor"

	"github.com/gorilla/websocket"
)

type WebSocketProcessor interface {
	processor.MessageProcessor
	SetConnection(connection *websocket.Conn)
}
