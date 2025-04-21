package websocket

import (
	"net/http"

	"github.com/gorilla/websocket"
)

// NewUpgrader создает и возвращает новый websocket.Upgrader с пользовательской
// функцией CheckOrigin. Функция CheckOrigin определяет, разрешен ли запрос
// на подключение websocket на основе заголовка Origin запроса.
//
// Параметры:
//   - debug: Если true, разрешены все origins.
//   - validOrigin: Конкретный origin, который разрешен, если debug равен false.
//
// Возвращает:
//
//	websocket.Upgrader, настроенный с пользовательской логикой CheckOrigin.
func NewUpgrader(debug bool, validOrigin string) websocket.Upgrader {
	return websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			if debug {
				return true
			}
			origin := r.Header.Get("Origin")
			return origin == validOrigin
		},
	}
}
