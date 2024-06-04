package http

import (
	"lazts/internal/utils"
)

const (
	DEFAULT_HOST = "localhost"
	DEFAULT_PORT = "8080"
)

type config struct {
	Host string
	Port string
}

func parseConfig() *config {
	host := utils.LookupEnv("HTTP_HOST", DEFAULT_HOST)
	port := utils.LookupNumericEnv("HTTP_PORT", DEFAULT_PORT)

	return &config{
		Host: host,
		Port: port,
	}
}
