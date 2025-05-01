package handlers

import (
	"fmt"
	"log"
	"messager/internal/ws/interfaces"
	"net/http"

	"github.com/gorilla/websocket"
)

type WebSocketHandler struct {
	upgrader         websocket.Upgrader
	messageSender    interfaces.WebSocketSender
	messageReceiver  interfaces.WebSocketReceiver
	messageProcessor interfaces.WebSocketProcessor
}

func NewWebSocketHandler(
	upgrader websocket.Upgrader,
	messageSender interfaces.WebSocketSender,
	messageReceiver interfaces.WebSocketReceiver,
	messageProcessor interfaces.WebSocketProcessor,
) *WebSocketHandler {
	return &WebSocketHandler{
		upgrader:         upgrader,
		messageSender:    messageSender,
		messageReceiver:  messageReceiver,
		messageProcessor: messageProcessor,
	}
}

// Tag возвращает строковый идентификатор для WebSocketHandler.
// Этот идентификатор может быть использован для логирования или отладки.
func (*WebSocketHandler) Tag() string {
	return "WEBSOCKET_HANDLER"
}

func (wsh *WebSocketHandler) HandleWebSocket(w http.ResponseWriter, r *http.Request) {

	conn, err := wsh.processConnection(w, r)
	if err != nil {
		http.Error(w, "Не удалось установить WebSocket соединение", http.StatusInternalServerError)
		log.Printf(wsh.Tag(), "Ошибка при апгрейде соединения:", err, "IP:", r.RemoteAddr)
		return
	}

	defer conn.Close()
	wsh.handleMessageLoop()
}

func (wsh *WebSocketHandler) handleMessageLoop() {
	for {
		message, err := wsh.messageReceiver.ReceiveMessage()
		if err != nil {
			wsh.handleError(err, "Ошибка чтения сообщения")
			break
		}

		responseMessage, err := wsh.messageProcessor.ProcessMessage(message)
		if err != nil {
			wsh.handleError(err, "Ошибка при обработке сообщения")
			break
		}

		if err := wsh.messageSender.SendResponseMessage(responseMessage); err != nil {
			wsh.handleError(err, "Ошибка при формировании ответа")
			break
		}
	}
}

// handleError обрабатывает ошибку, отправляя сообщение об ошибке с использованием
// отправителя сообщений службы WebSocket. Если отправка сообщения об ошибке не удалась,
// ошибка логируется. Кроме того, логируется исходная ошибка вместе с пользовательским сообщением.
//
// Параметры:
//   - err: Произошедшая ошибка.
//   - message: Пользовательское сообщение, описывающее контекст ошибки.
func (wsh *WebSocketHandler) handleError(err error, message string) {
	if err := wsh.messageSender.SendErrorMessage(message); err != nil {
		log.Println(wsh.Tag(), "Ошибка отправки сообщения об ошибке:", err)
	}
	log.Println(wsh.Tag(), message, err)
}

func (wsh *WebSocketHandler) processConnection(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	conn, err := wsh.upgrader.Upgrade(w, r, nil)
	if err != nil {
		return nil, fmt.Errorf("не удалось установить WebSocket соединение: %w", err)
	}

	wsh.messageSender.SetConnection(conn)
	wsh.messageReceiver.SetConnection(conn)
	wsh.messageProcessor.SetConnection(conn)

	return conn, nil
}
