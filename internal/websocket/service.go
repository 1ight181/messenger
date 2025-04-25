package websocket

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"messager/internal/websocket/interfaces"

	"os"

	"github.com/gorilla/websocket"
)

// WebsocketService представляет собой службу для обработки WebSocket соединений.
type WebsocketService struct {
	upgrader         websocket.Upgrader
	messageSender    interfaces.WebSocketSender
	messageReceiver  interfaces.WebSocketReceiver
	messageProcessor interfaces.WebSocketProcessor
}

// NewWebsocketService создает новый экземпляр WebsocketService с предоставленной
// конфигурацией и зависимостями.
//
// Параметры:
//   - debug: Логическое значение, указывающее, работает ли служба в режиме отладки.
//     Если true, служба будет разрешать подключения с любого источника.
//   - validOrigin: Строка, указывающая допустимый источник для WebSocket соединений.
//     Используется для проверки заголовка "Origin" во входящих запросах,
//     если режим отладки отключен.
//   - messageSender: Реализация интерфейса WebSocketSender, отвечающая за отправку
//     сообщений через WebSocket соединение.
//   - messageReceiver: Реализация интерфейса WebSocketReceiver, отвечающая за получение
//     сообщений из WebSocket соединения.
//   - messageProcessor: Реализация интерфейса WebSocketProcessor, отвечающая за обработку
//     сообщений, полученных из WebSocket соединения.
//
// Возвращает:
//
//	Указатель на экземпляр WebsocketService, настроенный с предоставленными параметрами.
func NewWebsocketService(
	debug bool,
	validOrigin string,
	messageSender interfaces.WebSocketSender,
	messageReceiver interfaces.WebSocketReceiver,
	messageProcessor interfaces.WebSocketProcessor,
) *WebsocketService {
	return &WebsocketService{
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				if debug {
					return true
				}
				origin := r.Header.Get("Origin")
				return origin == validOrigin
			},
		},
		messageSender:    messageSender,
		messageReceiver:  messageReceiver,
		messageProcessor: messageProcessor,
	}
}

// Tag возвращает строковый идентификатор для WebsocketService.
// Используется для представления службы с определенным тегом, "WEBSOCKET".
func (ws *WebsocketService) Tag() string {
	return "WEBSOCKET"
}

// handleWebSocket устанавливает WebSocket соединение и управляет его жизненным циклом.
// Оно обновляет HTTP соединение до WebSocket соединения, настраивает необходимые компоненты
// для отправки, получения и обработки сообщений, а также управляет циклом обмена сообщениями.
//
// Параметры:
//   - w: HTTP ответ для отправки данных клиенту.
//
// Поведение:
//   - Обновляет HTTP соединение до WebSocket соединения с использованием upgrader.
//   - Устанавливает WebSocket соединение для отправителя, получателя и обработчика сообщений.
//   - Непрерывно слушает входящие сообщения, обрабатывает их и отправляет ответы.
//   - Обрабатывает ошибки при получении, обработке или отправке сообщений и завершает
//     соединение в случае возникновения ошибки.
//
// Ошибки:
//   - Если обновление до WebSocket соединения не удалось, ошибка логируется, и отправляется
//     HTTP ответ с кодом 500.
//   - Если возникает ошибка при получении, обработке или отправке сообщений, ошибка
//     логируется, и соединение закрывается.
func (ws *WebsocketService) handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := ws.upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Не удалось установить WebSocket соединение", http.StatusInternalServerError)
		log.Printf(ws.Tag(), "Ошибка при апгрейде соединения:", err, "IP:", r.RemoteAddr)
		return
	}

	defer conn.Close()

	ws.messageSender.SetConnection(conn)
	ws.messageReceiver.SetConnection(conn)
	ws.messageProcessor.SetConnection(conn)

	for {
		message, err := ws.messageReceiver.ReceiveMessage()
		if err != nil {
			ws.handleError(err, "Ошибка чтения сообщения")
			break
		}

		responseMessage, err := ws.messageProcessor.ProcessMessage(message)
		if err != nil {
			ws.handleError(err, "Ошибка при обработке сообщения")
			break
		}

		if err := ws.messageSender.SendResponseMessage(responseMessage); err != nil {
			ws.handleError(err, "Ошибка при формировании ответа")
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
func (ws *WebsocketService) handleError(err error, message string) {
	if err := ws.messageSender.SendErrorMessage(message); err != nil {
		log.Println(ws.Tag(), "Ошибка отправки сообщения об ошибке:", err)
	}
	log.Println(ws.Tag(), message, err)
}

// StartServer запускает WebSocket сервер на указанном хосте и порту.
// Устанавливает HTTP обработчик для WebSocket эндпоинта на "/ws" и начинает
// прослушивание входящих соединений. Если сервер не удается запустить, ошибка
// логируется. Сервер корректно завершает работу при получении сигнала завершения.
//
// Параметры:
//   - host: Имя хоста или IP-адрес, на котором сервер будет слушать.
//   - port: Номер порта, на котором сервер будет слушать.
func (ws *WebsocketService) StartServer(host string, port string) {
	address := fmt.Sprintf("%s:%s", host, port)
	http.HandleFunc("/ws", ws.handleWebSocket)

	serverStopped := make(chan struct{})

	server := &http.Server{
		Addr:    address,
		Handler: nil,
	}

	go func() {
		log.Printf("Вебсокет запущен на: %s", address)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Ошибка запуска сервера:", err)
		}
		close(serverStopped)
	}()

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)

	select {
	case <-signalChannel:
		log.Println("Получен сигнал завершения, останавливаю сервер...")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			log.Println("Ошибка при остановке сервера:", err)
		}
	case <-serverStopped:
		log.Println("Сервер остановлен самостоятельно")
	}
}
