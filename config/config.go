package config

type Config struct {
	LogLevel   string
	DB         Database
	ListenPort int
}

type Database struct {
	DSN string
}
