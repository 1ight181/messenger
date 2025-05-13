package message

import "encoding/json"

type MessageType int

const (
	ErrorMessage MessageType = iota
	InfoMessage
	DataMessage

	ErrorResponse
	InfoResponse
	DataResponse
	UnknownResponse
)

// String возвращает строковое представление значения MessageType.
// Возвращаемое значение соответствует одному из предопределённых типов сообщений
// в зависимости от значения MessageType.
func (mt MessageType) String() string {
	return [...]string{"error", "info", "data", "unknown"}[mt]
}

// MarshalJSON реализует интерфейс json.Marshaler для типа MessageType.
// Он преобразует значение MessageType в строковое представление JSON,
// вызывая метод String и оборачивая результат в двойные кавычки.
func (mt MessageType) MarshalJSON() ([]byte, error) {
	return []byte(`"` + mt.String() + `"`), nil
}

// UnmarshalJSON реализует пользовательский JSON-демаршалер для типа MessageType.
// Он интерпретирует JSON-данные как строку и сопоставляет её с соответствующим
// значением MessageType. Если строка не соответствует ни одному известному
// значению MessageType, присваивается недопустимое значение (-1).
//
// Параметры:
//   - b: Срез байтов, содержащий JSON-данные.
//
// Возвращает:
//   - Ошибку, если входные данные не могут быть демаршалированы в строку, или nil,
//     если демаршалирование и сопоставление выполнены успешно.
func (mt *MessageType) UnmarshalJSON(b []byte) error {
	var str string
	if err := json.Unmarshal(b, &str); err != nil {
		return err
	}

	switch str {
	case "error":
		*mt = ErrorMessage
	case "info":
		*mt = InfoMessage
	case "data":
		*mt = DataMessage
	default:
		*mt = -1 // Неизвестный тип
	}
	return nil
}
