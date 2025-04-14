package websocket

import (
	"net/http"

	"github.com/gorilla/websocket"
)

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
