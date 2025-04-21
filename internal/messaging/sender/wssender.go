package sender

import (
	models "messager/internal/messaging/models"

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

func (wsms *WebSocketMessageSender) SendMessage(message models.Message) error {
	return wsms.connection.WriteJSON(message)
}

func (wsms *WebSocketMessageSender) SendErrorMessage(messageText string) error {
	errorMessage := models.Message{
		Type: "error",
		Text: messageText,
	}
	return wsms.SendMessage(errorMessage)
}

func (wsms *WebSocketMessageSender) SendInfoMessage(messageText string) error {
	infoMessage := models.Message{
		Type: "info",
		Text: messageText,
	}
	return wsms.SendMessage(infoMessage)
}

func (wsms *WebSocketMessageSender) SendDataMessage(messageText string) error {
	infoMessage := models.Message{
		Type: "data",
		Text: messageText,
	}
	return wsms.SendMessage(infoMessage)
}

func (wsms *WebSocketMessageSender) SendResponseMessage(messageText models.Message) error {
	return wsms.SendMessage(messageText)
}
