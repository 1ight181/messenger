package processor

import (
	models "messager/internal/messaging/models"
)

type MessageProcessor interface {
	ProcessMessage(models.Message) (models.Message, error) //обработать сообщение и вернуть ответное сообщение
}
