package processor

import (
	models "messenger/internal/messaging/models"
)

type MessageProcessor interface {
	ProcessMessage(models.Message) (models.Message, error) //обработать сообщение и вернуть ответное сообщение
}
