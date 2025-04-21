package config_types

type WebSocket struct {
	Host  string `mapstructure:"host"`
	Port  string `mapstructure:"port"`
	Debug bool   `mapstructure:"debug"`
}
