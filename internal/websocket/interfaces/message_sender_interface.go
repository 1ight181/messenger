package interfaces

import (
	"messager/internal/messaging/sender"
	types "messager/internal/messaging/types"

	"github.com/gorilla/websocket"
)

type WebSocketSender interface {
	sender.MessageSender
	SendErrorMessage(message string) error
	SendInfoMessage(message string) error
	SendResponseMessage(message types.Message) error
	SetConnection(connection *websocket.Conn)
}
