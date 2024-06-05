package markdown

import (
	"lazts/internal/utils"
)

const (
	DEFAULT_THEME           = "dracula"
	DEFAULT_WORD_PER_MINUTE = 250
)

type config struct {
	Theme         string
	WordPerMinute int
}

func parseConfig() *config {
	return &config{
		Theme:         utils.LookupEnv("MARKDOWN_THEME", DEFAULT_THEME),
		WordPerMinute: int(utils.LookupUIntEnv("MARKDOWN_WORD_PER_MINUTE", DEFAULT_WORD_PER_MINUTE)),
	}
}
