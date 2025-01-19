package config

type Config struct {
	Server   string
	Port     int
	Database string
}

type ConfigProvider interface {
	GetConfig() Config
}
