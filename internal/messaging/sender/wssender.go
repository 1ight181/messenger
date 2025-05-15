package sender

import (
	"errors"
	msg "messenger/internal/messaging/models/message"
	"time"

	"github.com/gorilla/websocket"
)

type WebSocketMessageSender struct {
	connection *websocket.Conn
}

type Options struct {
	// Здесь можно добавить дополнительные параметры конфигурации
}

// New создает и возвращает новый экземпляр WebSocketMessageSender, используя предоставленные Options.
// Возвращаемый отправитель инициализируется с указанными параметрами конфигурации.
func New(options Options) *WebSocketMessageSender {
	return &WebSocketMessageSender{
		connection: nil,
		// option: options.option,
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
	if wsms.connection == nil {
		return errors.New("соединение не установлено")
	}
	return wsms.connection.WriteJSON(message)
}

func (wsms *WebSocketMessageSender) SendCloseMessage(code int, text string, timeout time.Duration) error {
	if wsms.connection == nil {
		return errors.New("соединение не установлено")
	}

	return wsms.connection.WriteControl(
		websocket.CloseMessage,
		websocket.FormatCloseMessage(code, text),
		time.Now().Add(timeout),
	)
}
