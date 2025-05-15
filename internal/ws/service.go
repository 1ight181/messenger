package ws

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"os"
)

// WebsocketService представляет собой службу для обработки WebSocket соединений.
type WebsocketService struct {
	server *http.Server
}

// NewWebsocketService создает новый экземпляр WebsocketService.
// Принимает HTTP сервер в качестве параметра и возвращает указатель на инициализированный WebsocketService.
//
// Параметры:
//   - server: Экземпляр *http.Server, который будет использоваться WebsocketService.
//
// Возвращает:
//   - Указатель на экземпляр WebsocketService.
func NewWebsocketService(
	server *http.Server,
) *WebsocketService {
	return &WebsocketService{
		server: server,
	}
}

// Tag возвращает строковый идентификатор для WebsocketService.
// Используется для представления службы с определенным тегом, "WEBSOCKET".
func (ws *WebsocketService) Tag() string {
	return "WEBSOCKET"
}

// StartServer запускает веб-сервер с поддержкой WebSocket и обрабатывает его завершение по сигналу остановки.
// Сервер запускается в отдельной горутине и слушает указанный адрес с использованием TLS.
// При получении сигнала завершения (например, SIGTERM или прерывания) инициируется корректное завершение работы сервера с таймаутом.
// В случае ошибок при запуске или остановке сервера выводятся соответствующие сообщения в лог.
func (ws *WebsocketService) StartServer() {
	stopSignal := make(chan os.Signal, 1)
	signal.Notify(stopSignal, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Printf("Вебсокет запущен на: %s", ws.server.Addr)
		if err := ws.server.ListenAndServeTLS("", ""); err != nil && err != http.ErrServerClosed {
			log.Fatal("Ошибка запуска сервера:", err)
		}
	}()

	<-stopSignal
	log.Println("Получен сигнал завершения, сервер останавливается...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := ws.server.Shutdown(shutdownCtx); err != nil {
		log.Println("Ошибка при остановке сервера:", err)
	}
}
