package receiver

import (
	"log"
	models "messager/internal/messaging/models"

	"github.com/gorilla/websocket"
)

type WebSocketMessageReceiver struct {
	connection *websocket.Conn
}

func NewWebSocketMessageReceiver() *WebSocketMessageReceiver {
	return &WebSocketMessageReceiver{}
}

func (wsmr *WebSocketMessageReceiver) SetConnection(conn *websocket.Conn) {
	wsmr.connection = conn
}

func (wsmr *WebSocketMessageReceiver) ReceiveMessage() (models.Message, error) {
	if wsmr.connection != nil {
		var message models.Message
		if err := wsmr.connection.ReadJSON(&message); err != nil {

			return models.Message{}, err
		}
		log.Printf("Получено сообщение: %s\n", message)
		return message, nil
	} else {
		return models.Message{}, &websocket.CloseError{Code: websocket.CloseAbnormalClosure, Text: "Connection is not set"}
	}

}
