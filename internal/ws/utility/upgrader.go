package utility

import (
	"net/http"
	"strings"

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
func NewUpgrader(debug bool, invalidOrigins []string) websocket.Upgrader {
	return websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			if debug {
				return true
			}
			origin := r.Header.Get("Origin")
			return isAllowedOrigin(origin, invalidOrigins)
		},
	}
}

// isAllowedOrigin проверяет, разрешен ли данный origin на основе списка запрещенных origin'ов.
// Возвращает false, если origin пустой или содержит любой из запрещенных подстрок,
// указанных в срезе invalidOrigins. Если срез invalidOrigins пуст, все origin'ы считаются разрешенными.
//
// Параметры:
//   - origin: Строка origin для проверки.
//   - invalidOrigins: Срез строк, представляющих подстроки origin'ов, которые запрещены.
//
// Возвращает:
//   - bool: True, если origin разрешен, false в противном случае.
func isAllowedOrigin(origin string, invalidOrigins []string) bool {
	if len(origin) == 0 {
		return false
	}

	if len(invalidOrigins) == 0 {
		return true
	}

	for _, blocked := range invalidOrigins {
		if strings.Contains(origin, blocked) {
			return false
		}
	}

	return true
}
