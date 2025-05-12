package interfaces

import (
	models "messenger/internal/messaging/models"
)

type MessageSender interface {
	SendMessage(message models.Message) error
}
