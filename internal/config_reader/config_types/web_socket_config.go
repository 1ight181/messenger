package config_types

type WebSocketConfig struct {
	Host  string `mapstructure:"host"`
	Port  string `mapstructure:"port"`
	Debug bool   `mapstructure:"debug"`
}
