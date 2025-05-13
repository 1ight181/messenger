package sender

import (
	msg "messenger/internal/messaging/models/message"

	"github.com/gorilla/websocket"
)

type WebSocketMessageSender struct {
	connection *websocket.Conn
}

// NewWebSocketMessageSender создает и возвращает новый экземпляр WebSocketMessageSender.
// Эта функция инициализирует WebSocketMessageSender с значениями по умолчанию.
func New() *WebSocketMessageSender {
	return &WebSocketMessageSender{
		connection: nil,
	}
}

// SetConnection устанавливает WebSocket-соединение для WebSocketMessageSender.
// Этот метод присваивает предоставленный экземпляр websocket.Conn внутреннему полю соединения отправителя.
//
// Параметры:
//   - conn: Указатель на websocket.Conn, представляющий WebSocket-соединение, которое будет использоваться.
func (wsms *WebSocketMessageSender) SetConnection(conn *websocket.Conn) {
	wsms.connection = conn
}

// SendMessage отправляет сообщение через WebSocket-соединение.
// Принимает msg.Message в качестве входного параметра и записывает его в формате JSON в WebSocket.
// Возвращает ошибку, если сообщение не может быть отправлено или если возникли проблемы с соединением.
func (wsms *WebSocketMessageSender) SendMessage(message msg.Message) error {
	return wsms.connection.WriteJSON(message)
}
