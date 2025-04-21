package websocket

import (
	"fmt"
	"log"
	"net/http"

	"messager/internal/websocket/interfaces"

	"github.com/gorilla/websocket"
)

type WebsocketService struct {
	upgrader         websocket.Upgrader
	messageSender    interfaces.WebSocketSender
	messageReceiver  interfaces.WebSocketReceiver
	messageProcessor interfaces.WebSocketProcessor
}

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

func (ws *WebsocketService) Tag() string {
	return "WEBSOCKET"
}

func (ws *WebsocketService) handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := ws.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatalln(ws.Tag(), "Ошибка при апргрейде соедиенения:", err)
		http.Error(w, "Не удалось установить WebSocket соединение", http.StatusInternalServerError)
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
			ws.handleError(err, "Ошибка чтения сообщения")
			break
		}

		if err := ws.messageSender.SendResponseMessage(responseMessage); err != nil {
			ws.handleError(err, "Ошибка при формировании ответа")
			break
		}
	}
}

func (ws *WebsocketService) handleError(err error, message string) {
	if err := ws.messageSender.SendErrorMessage(message); err != nil {
		log.Println(ws.Tag(), "Ошибка отправки сообщения об ошибке:", err)
	}
	log.Println(ws.Tag(), message, err)
}

func (ws *WebsocketService) StartServer(host string, port string) {
	address := fmt.Sprintf("%s:%s", host, port)
	http.HandleFunc("/ws", ws.handleWebSocket)

	log.Printf("Вебсокет запущен на: %s", address)
	if err := http.ListenAndServe(address, nil); err != nil {
		log.Fatal("Ошибка запуска сервера:", err)
	}
}
