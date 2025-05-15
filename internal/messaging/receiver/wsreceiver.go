package receiver

import (
	"log"
	msg "messenger/internal/messaging/models/message"

	"github.com/gorilla/websocket"
)

type WebSocketMessageReceiver struct {
	connection *websocket.Conn
}

type Options struct {
	// Здесь можно добавить дополнительные параметры конфигурации
}

func New(options Options) *WebSocketMessageReceiver {
	return &WebSocketMessageReceiver{
		connection: nil,
		// option: options.option,
	}
}

// SetConnection устанавливает WebSocket-соединение для WebSocketMessageReceiver.
// Этот метод присваивает переданный экземпляр websocket.Conn в поле connection получателя.
//
// Параметры:
//   - conn: Указатель на экземпляр websocket.Conn, представляющий WebSocket-соединение.
func (wsmr *WebSocketMessageReceiver) SetConnection(conn *websocket.Conn) {
	wsmr.connection = conn
}

// ReceiveMessage читает сообщение в формате JSON из WebSocket-соединения
// и возвращает его как экземпляр msg.Message. Если соединение не установлено,
// возвращается ошибка websocket.CloseError, указывающая на ненормальное закрытие.
// В случае ошибки при чтении возвращается ошибка вместе с пустым сообщением.
//
// Возвращает:
// - msg.Message: Декодированное сообщение из WebSocket-соединения.
// - error: Ошибка, если соединение не установлено или если чтение сообщения завершилось неудачей.
func (wsmr *WebSocketMessageReceiver) ReceiveMessage() (msg.Message, error) {
	if wsmr.connection != nil {
		var message msg.Message
		if err := wsmr.connection.ReadJSON(&message); err != nil {
			return msg.Message{}, err
		}
		log.Printf("Получено сообщение: %s\n", message)
		return message, nil
	} else {
		return msg.Message{}, &websocket.CloseError{Code: websocket.CloseAbnormalClosure, Text: "Connection is not set"}
	}
}
