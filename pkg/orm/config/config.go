package config

type Config struct {
    SQLDir        string
    DriverName    string
    DSN           string
    MaxOpenConns  int
    MaxIdleConns  int
}
