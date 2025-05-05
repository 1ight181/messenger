package receiver

import (
	"log"
	models "messenger/internal/messaging/models"

	"github.com/gorilla/websocket"
)

type WebSocketMessageReceiver struct {
	connection *websocket.Conn
}

// NewWebSocketMessageReceiver создает и возвращает новый экземпляр WebSocketMessageReceiver.
// Эта функция инициализирует WebSocketMessageReceiver с настройками по умолчанию.
func NewWebSocketMessageReceiver() *WebSocketMessageReceiver {
	return &WebSocketMessageReceiver{}
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
// и возвращает его как экземпляр models.Message. Если соединение не установлено,
// возвращается ошибка websocket.CloseError, указывающая на ненормальное закрытие.
// В случае ошибки при чтении возвращается ошибка вместе с пустым сообщением.
//
// Возвращает:
// - models.Message: Декодированное сообщение из WebSocket-соединения.
// - error: Ошибка, если соединение не установлено или если чтение сообщения завершилось неудачей.
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
