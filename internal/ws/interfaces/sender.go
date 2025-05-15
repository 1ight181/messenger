package interfaces

import (
	"messenger/internal/messaging/interfaces"
	"time"

	"github.com/gorilla/websocket"
)

type WebSocketSender interface {
	interfaces.MessageSender
	SetConnection(connection *websocket.Conn)
	SendCloseMessage(code int, text string, timeout time.Duration) error
}
