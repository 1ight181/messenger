package processor

import (
	"log"
	messaging "messager/internal/messaging/types"

	"github.com/gorilla/websocket"
)

type WebSocketMessageProcessor struct {
	connection *websocket.Conn
}

func NewWebSocketMessageProcessor() *WebSocketMessageProcessor {
	return &WebSocketMessageProcessor{}
}

func (wsmp *WebSocketMessageProcessor) SetConnection(conn *websocket.Conn) {
	wsmp.connection = conn
}

func (wsmp *WebSocketMessageProcessor) ProcessMessage(message messaging.Message) (messaging.Message, error) {
	if wsmp.connection != nil {
		switch message.Type {
		case "error":
			responseMessage := wsmp.processError(message)
			return responseMessage, nil
		case "info":
			responseMessage := wsmp.processInfo(message)
			return responseMessage, nil
		case "data":
			responseMessage := wsmp.processData(message)
			return responseMessage, nil
		default:
			log.Printf("Получен неизвестный тип сообщения: %s\n", message)
			responseMessage := messaging.Message{
				Type: "error_unknown_message_type",
				Text: "Неизвестный тип сообщения",
			}
			return responseMessage, nil
		}
	} else {
		return messaging.Message{}, &websocket.CloseError{Code: websocket.CloseAbnormalClosure, Text: "Connection is not set"}
	}
}

func (wsmp *WebSocketMessageProcessor) processError(errorMessage messaging.Message) messaging.Message {
	log.Printf("Клиент отправил сообщение об ошибке: %s\n", errorMessage.Text)
	responseMessage := messaging.Message{
		Type: "error_response",
		Text: "Сообщение об ошибке получено",
	}
	return responseMessage
}

func (wsmp *WebSocketMessageProcessor) processInfo(infoMessage messaging.Message) messaging.Message {
	log.Printf("Клиент отправил информационное сообщение: %s\n", infoMessage.Text)
	responseMessage := messaging.Message{
		Type: "info_response",
		Text: "Информационное сообщение получено",
	}
	return responseMessage
}

func (wsmp *WebSocketMessageProcessor) processData(dataMessage messaging.Message) messaging.Message {
	log.Printf("Клиент отправил сообщение с данными: %s\n", dataMessage.Text)
	responseMessage := messaging.Message{
		Type: "data_response",
		Text: "Сообщение с данными получено",
	}
	return responseMessage
}
