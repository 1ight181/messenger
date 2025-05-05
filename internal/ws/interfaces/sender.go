package interfaces

import (
	models "messenger/internal/messaging/models"
	"messenger/internal/messaging/sender"

	"github.com/gorilla/websocket"
)

type WebSocketSender interface {
	sender.MessageSender
	SendErrorMessage(message string) error
	SendInfoMessage(message string) error
	SendResponseMessage(message models.Message) error
	SetConnection(connection *websocket.Conn)
}
