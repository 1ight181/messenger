package config_types

type Config struct {
	WebSocket WebSocketConfig `mapstructure:"ws"`
}

func (c *Config) GetWebSocketHost() string {
	return c.WebSocket.Host
}

func (c *Config) GetWebSocketPort() int {
	return c.WebSocket.Port
}
