package markdown

import (
	"lazts/internal/utils"
)

const (
	DEFAULT_THEME = "dracula"
)

type config struct {
	Theme string
}

func parseConfig() *config {
	return &config{
		Theme: utils.LookupEnv("MARKDOWN_THEME", DEFAULT_THEME),
	}
}
