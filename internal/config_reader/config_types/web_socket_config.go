package config_types

type WebSocketConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}
