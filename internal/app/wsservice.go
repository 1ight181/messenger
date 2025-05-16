package app

import (
	"crypto/tls"
	"fmt"

	"messenger/internal/config/loaders"
	"messenger/internal/config/models"
	wshfac "messenger/internal/factories/wshandler"

	processor "messenger/internal/messaging/processor"
	receiver "messenger/internal/messaging/receiver"
	sender "messenger/internal/messaging/sender"

	ws "messenger/internal/ws"
	wsupgr "messenger/internal/ws/upgraders"
	"net/http"
)

type WebSocketServiceOptions struct {
	Config           models.WebSocket
	TLSConfig        *tls.Config
	SenderOptions    sender.Options
	ReceiverOptions  receiver.Options
	ProcessorOptions processor.Options
}

func loadAppWebSocketService(opts WebSocketServiceOptions) *ws.WebsocketService {
	wsHost, wsPort, wsDebug, invalidOrigins := loaders.LoadWebsocket(opts.Config)

	upgrager := wsupgr.NewUpgrader(wsDebug, invalidOrigins)

	handlerFactoryOptions := wshfac.Options{
		Upgrader:         upgrager,
		SenderOptions:    opts.SenderOptions,
		ReceiverOptions:  opts.ReceiverOptions,
		ProcessorOptions: opts.ProcessorOptions,
	}

	webSocketHandlerFactory := wshfac.New(handlerFactoryOptions)

	wsHandlerFunc := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler := webSocketHandlerFactory.NewHandler()
		handler.HandleWebSocket(w, r)
	})

	address := fmt.Sprintf("%s:%s", wsHost, wsPort)

	httpServer := &http.Server{
		Addr:      address,
		Handler:   wsHandlerFunc,
		TLSConfig: opts.TLSConfig,
	}

	return ws.NewWebsocketService(httpServer)
}
