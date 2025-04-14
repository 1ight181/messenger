package receiver

import (
	"log"
	types "messager/internal/messaging/types"

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

func (wsmr *WebSocketMessageReceiver) ReceiveMessage() (types.Message, error) {
	if wsmr.connection != nil {
		var message types.Message
		if err := wsmr.connection.ReadJSON(&message); err != nil {

			return types.Message{}, err
		}
		log.Printf("Получено сообщение: %s\n", message)
		return message, nil
	} else {
		return types.Message{}, &websocket.CloseError{Code: websocket.CloseAbnormalClosure, Text: "Connection is not set"}
	}

}
