package interfaces

import (
	"messenger/internal/messaging/interfaces"
	models "messenger/internal/messaging/models"

	"github.com/gorilla/websocket"
)

type WebSocketSender interface {
	interfaces.MessageSender
	SendErrorMessage(message string) error
	SendInfoMessage(message string) error
	SendResponseMessage(message models.Message) error
	SetConnection(connection *websocket.Conn)
}
