package processor

import (
	"log"
	models "messenger/internal/messaging/models"

	"github.com/gorilla/websocket"
)

// WebSocketMessageProcessor представляет процессор для обработки сообщений WebSocket.
type WebSocketMessageProcessor struct {
	connection *websocket.Conn
}

// NewWebSocketMessageProcessor создает и возвращает новый экземпляр WebSocketMessageProcessor.
// Эта функция инициализирует WebSocketMessageProcessor с настройками по умолчанию.
func NewWebSocketMessageProcessor() *WebSocketMessageProcessor {
	return &WebSocketMessageProcessor{}
}

// SetConnection устанавливает WebSocket-соединение для WebSocketMessageProcessor.
// Этот метод присваивает предоставленный экземпляр websocket.Conn в поле connection процессора.
//
// Параметры:
//   - conn: Указатель на websocket.Conn, представляющий устанавливаемое WebSocket-соединение.
func (wsmp *WebSocketMessageProcessor) SetConnection(conn *websocket.Conn) {
	wsmp.connection = conn
}

// ProcessMessage обрабатывает входящее сообщение WebSocket в зависимости от его типа.
// Он обрабатывает различные типы сообщений ("error", "info", "data") с помощью
// соответствующих методов обработки и возвращает ответное сообщение.
//
// Параметры:
//   - message: Входящее сообщение типа models.Message, которое нужно обработать.
//
// Возвращает:
//   - models.Message: Обработанное ответное сообщение.
//   - error: Ошибка, если WebSocket-соединение не установлено или произошла другая проблема.
//
// Поведение:
//   - Если WebSocket-соединение не установлено, возвращает ошибку типа *websocket.CloseError.
//   - Для известных типов сообщений ("error", "info", "data") обрабатывает сообщение с помощью
//     соответствующих методов (processError, processInfo, processData).
//   - Для неизвестных типов сообщений регистрирует проблему и возвращает ответное сообщение
//     с типом "error_unknown_message_type" и описанием ошибки.
func (wsmp *WebSocketMessageProcessor) ProcessMessage(message models.Message) (models.Message, error) {
	if wsmp.connection != nil {
		switch message.Type {
		case "error":
			responseMessage := wsmp.processError(message)
			return responseMessage, nil
		case "info":
			responseMessage := wsmp.processInfo(message)
			return responseMessage, nil
		case "data":
			responseMessage := wsmp.processData(message)
			return responseMessage, nil
		default:
			log.Printf("Получен неизвестный тип сообщения: %s\n", message)
			responseMessage := models.Message{
				Type: "error_unknown_message_type",
				Text: "Неизвестный тип сообщения",
			}
			return responseMessage, nil
		}
	} else {
		return models.Message{}, &websocket.CloseError{Code: websocket.CloseAbnormalClosure, Text: "Connection is not set"}
	}
}

// processError обрабатывает сообщение об ошибке, полученное от клиента.
// Он регистрирует сообщение об ошибке и возвращает ответное сообщение,
// указывающее, что сообщение об ошибке было получено.
//
// Параметры:
//   - errorMessage: Сообщение об ошибке, отправленное клиентом.
//
// Возвращает:
//   - Объект models.Message, содержащий ответное сообщение с типом "error_response"
//     и текстом, указывающим, что сообщение об ошибке было получено.
func (wsmp *WebSocketMessageProcessor) processError(errorMessage models.Message) models.Message {
	log.Printf("Клиент отправил сообщение об ошибке: %s\n", errorMessage.Text)
	responseMessage := models.Message{
		Type: "error_response",
		Text: "Сообщение об ошибке получено",
	}
	return responseMessage
}

// processInfo обрабатывает информационное сообщение, полученное от клиента.
// Он регистрирует содержимое сообщения и возвращает ответное сообщение,
// подтверждающее получение информационного сообщения.
//
// Параметры:
//   - infoMessage (models.Message): Информационное сообщение, отправленное клиентом.
//
// Возвращает:
//   - models.Message: Ответное сообщение, указывающее, что информационное
//     сообщение было получено.
func (wsmp *WebSocketMessageProcessor) processInfo(infoMessage models.Message) models.Message {
	log.Printf("Клиент отправил информационное сообщение: %s\n", infoMessage.Text)
	responseMessage := models.Message{
		Type: "info_response",
		Text: "Информационное сообщение получено",
	}
	return responseMessage
}

// processData обрабатывает входящее сообщение с данными и генерирует ответное сообщение.
// Он регистрирует текст полученного сообщения и возвращает предопределенный ответ.
//
// Параметры:
//   - dataMessage: Входящее сообщение типа models.Message, содержащее данные.
//
// Возвращает:
//   - models.Message: Ответное сообщение с типом "data_response" и предопределенным текстом.
func (wsmp *WebSocketMessageProcessor) processData(dataMessage models.Message) models.Message {
	log.Printf("Клиент отправил сообщение с данными: %s\n", dataMessage.Text)
	responseMessage := models.Message{
		Type: "data_response",
		Text: "Сообщение с данными получено",
	}
	return responseMessage
}
