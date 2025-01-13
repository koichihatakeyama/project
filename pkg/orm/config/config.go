package config

type Config struct {
	SQLDir       string
	DriverName   string
	DSN          string
	MaxOpenConns int
	MaxIdleConns int
	QueryTimeout int
}

func NewConfig() *Config {
	return &Config{
		MaxOpenConns: 10,
		MaxIdleConns: 5,
		QueryTimeout: 30,
	}
}
