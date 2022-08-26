package config

const (
	defaultPort = 8080
	message     = "Default"
)

type Config struct {
	Port    int
	message string
}
