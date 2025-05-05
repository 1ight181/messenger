package sender

import (
	models "messenger/internal/messaging/models"

	"github.com/gorilla/websocket"
)

type WebSocketMessageSender struct {
	connection *websocket.Conn
}

// NewWebSocketMessageSender создает и возвращает новый экземпляр WebSocketMessageSender.
// Эта функция инициализирует WebSocketMessageSender с значениями по умолчанию.
func NewWebSocketMessageSender() *WebSocketMessageSender {
	return &WebSocketMessageSender{}
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
// Принимает models.Message в качестве входного параметра и записывает его в формате JSON в WebSocket.
// Возвращает ошибку, если сообщение не может быть отправлено или если возникли проблемы с соединением.
func (wsms *WebSocketMessageSender) SendMessage(message models.Message) error {
	return wsms.connection.WriteJSON(message)
}

// SendErrorMessage отправляет сообщение об ошибке с указанным текстом, используя WebSocketMessageSender.
// Сообщение создается с типом "error" и предоставленным текстом.
// Возвращает ошибку, если сообщение не удалось отправить.
//
// Параметры:
//   - messageText: Текстовое содержимое сообщения об ошибке.
//
// Возвращает:
//   - error: Ошибка, если сообщение не удалось отправить, или nil, если отправка успешна.
func (wsms *WebSocketMessageSender) SendErrorMessage(messageText string) error {
	errorMessage := models.Message{
		Type: "error",
		Text: messageText,
	}
	return wsms.SendMessage(errorMessage)
}

// SendInfoMessage отправляет информационное сообщение через WebSocket-соединение.
// Сообщение создается с предоставленным текстом и типом "info".
// Использует метод SendMessage для передачи сообщения.
//
// Параметры:
//   - messageText: Текстовое содержимое информационного сообщения.
//
// Возвращает:
//   - error: Ошибка, если сообщение не удалось отправить, или nil, если отправка успешна.
func (wsms *WebSocketMessageSender) SendInfoMessage(messageText string) error {
	infoMessage := models.Message{
		Type: "info",
		Text: messageText,
	}
	return wsms.SendMessage(infoMessage)
}

// SendDataMessage отправляет сообщение с данными с указанным текстом через WebSocket-соединение.
// Сообщение создается с типом "data" и предоставленным текстом, а затем передается
// методу SendMessage для отправки.
//
// Параметры:
//   - messageText: Текстовое содержимое сообщения с данными.
//
// Возвращает:
//   - error: Ошибка, если сообщение не удалось отправить, или nil, если отправка успешна.
func (wsms *WebSocketMessageSender) SendDataMessage(messageText string) error {
	infoMessage := models.Message{
		Type: "data",
		Text: messageText,
	}
	return wsms.SendMessage(infoMessage)
}

// SendResponseMessage отправляет ответное сообщение, используя WebSocketMessageSender.
// Принимает models.Message в качестве входного параметра и передает его методу SendMessage.
// Возвращает ошибку, если сообщение не удалось отправить.
func (wsms *WebSocketMessageSender) SendResponseMessage(messageText models.Message) error {
	return wsms.SendMessage(messageText)
}
