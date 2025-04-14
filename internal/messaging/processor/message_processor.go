package processor

import (
	messaging "messager/internal/messaging/types"
)

type MessageProcessor interface {
	ProcessMessage(messaging.Message) (messaging.Message, error) //обработать сообщение и вернуть ответное сообщение
}
