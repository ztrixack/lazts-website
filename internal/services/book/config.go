package book

import (
	"lazts/internal/utils"
)

const (
	DEFAULT_CONTENT_DIR = "./web/contents"
)

type config struct {
	ContentDir string
}

func parseConfig() *config {
	return &config{
		ContentDir: utils.LookupEnv("CONTENT_DIR", DEFAULT_CONTENT_DIR),
	}
}
