package handlers

import (
	"fmt"
	"log"
	msg "messenger/internal/messaging/models/message"
	"messenger/internal/ws/interfaces"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type WebSocketHandler struct {
	upgrader         websocket.Upgrader
	messageSender    interfaces.WebSocketSender
	messageReceiver  interfaces.WebSocketReceiver
	messageProcessor interfaces.WebSocketProcessor
}

func New(
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

// HandleWebSocket устанавливает WebSocket соединение и обрабатывает входящие сообщения.
// Он апгрейдит HTTP соединение до WebSocket соединения и запускает цикл обработки сообщений.
//
// Параметры:
//   - w: HTTP ответ для отправки ответов клиенту.
//   - r: HTTP запрос, содержащий запрос на апгрейд до WebSocket.
//
// Поведение:
//   - Пытается апгрейдить HTTP соединение до WebSocket соединения.
//   - Если апгрейд не удался, возвращает ошибку HTTP 500 и логирует детали ошибки.
//   - Если апгрейд успешен, запускает цикл обработки сообщений и гарантирует закрытие соединения по завершении.
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

// handleMessageLoop выполняет непрерывную обработку входящих WebSocket сообщений в цикле.
// Он выполняет следующие шаги:
// 1. Получает сообщение с использованием messageReceiver.
// 2. Обрабатывает полученное сообщение с использованием messageProcessor.
// 3. Отправляет обработанное сообщение-ответ с использованием messageSender.
//
// Если на любом этапе (получение, обработка или отправка) возникает ошибка,
// метод обрабатывает её с помощью handleError и завершает цикл.
//
// Этот метод предназначен для работы до тех пор, пока не произойдет ошибка.
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

		if err := wsh.messageSender.SendMessage(responseMessage); err != nil {
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
	errorResponse := msg.NewErrorMessage("Ошибка отправки сообщения об ошибке")

	if websocket.IsCloseError(err,
		websocket.CloseNormalClosure,
		websocket.CloseGoingAway,
		websocket.CloseAbnormalClosure,
	) {
		wsh.handleConnectionClose(err)
		return
	}

	if err := wsh.messageSender.SendMessage(errorResponse); err != nil {
		log.Println(wsh.Tag(), "Ошибка отправки сообщения об ошибке:", err)
	}
	log.Println(wsh.Tag(), message, err)
}

func (wsh *WebSocketHandler) handleConnectionClose(err error) {
	if closeErr, ok := err.(*websocket.CloseError); ok {
		log.Printf("Соединение закрыто: код %d, причина: %s\n", closeErr.Code, closeErr.Text)
		wsh.messageSender.SendCloseMessage(closeErr.Code, "Закрытие обработано", 3*time.Second)

	} else {
		log.Println(wsh.Tag(), "Соединение закрыто:", err)
	}
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
