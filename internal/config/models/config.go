package models

type Config struct {
	WebSocket   WebSocket   `mapstructure:"ws"`
	Certificate Certificate `mapstructure:"certificate"`
}

// Validate проверяет поля конфигурации структуры Config на корректность.
// Она проверяет конфигурации WebSocket и Certificate, вызывая их
// соответствующие методы Validate. Если какая-либо проверка не проходит,
// возвращается ошибка; в противном случае возвращается nil.
func (c *Config) Validate() error {
	if err := c.WebSocket.Validate(); err != nil {
		return err
	}
	if err := c.Certificate.Validate(); err != nil {
		return err
	}
	return nil
}
