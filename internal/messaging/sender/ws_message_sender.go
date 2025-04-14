package sender

import (
	types "messager/internal/messaging/types"

	"github.com/gorilla/websocket"
)

type WebSocketMessageSender struct {
	connection *websocket.Conn
}

func NewWebSocketMessageSender() *WebSocketMessageSender {
	return &WebSocketMessageSender{}
}

func (wsms *WebSocketMessageSender) SetConnection(conn *websocket.Conn) {
	wsms.connection = conn
}

func (wsms *WebSocketMessageSender) SendMessage(message types.Message) error {
	return wsms.connection.WriteJSON(message)
}

func (wsms *WebSocketMessageSender) SendErrorMessage(messageText string) error {
	errorMessage := types.Message{
		Type: "error",
		Text: messageText,
	}
	return wsms.SendMessage(errorMessage)
}

func (wsms *WebSocketMessageSender) SendInfoMessage(messageText string) error {
	infoMessage := types.Message{
		Type: "info",
		Text: messageText,
	}
	return wsms.SendMessage(infoMessage)
}

func (wsms *WebSocketMessageSender) SendDataMessage(messageText string) error {
	infoMessage := types.Message{
		Type: "data",
		Text: messageText,
	}
	return wsms.SendMessage(infoMessage)
}

func (wsms *WebSocketMessageSender) SendResponseMessage(messageText types.Message) error {
	return wsms.SendMessage(messageText)
}
