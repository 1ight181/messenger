package sender

import (
	models "messager/internal/messaging/models"
)

type MessageSender interface {
	SendMessage(message models.Message) error
}
