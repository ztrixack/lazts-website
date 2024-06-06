package markdown

import (
	"lazts/internal/utils"
)

const (
	DEFAULT_THEME           = "dracula"
	DEFAULT_WORD_PER_MINUTE = 250
	DEFAULT_CONTENT_DIR     = "./web/contents"
	DEFAULT_CONTENT_FILE    = "index.md"
)

type config struct {
	Theme         string
	WordPerMinute int
	ContentDir    string
	ContentFile   string
}

func parseConfig() *config {
	return &config{
		Theme:         utils.LookupEnv("MARKDOWN_THEME", DEFAULT_THEME),
		WordPerMinute: int(utils.LookupUIntEnv("MARKDOWN_WORD_PER_MINUTE", DEFAULT_WORD_PER_MINUTE)),
		ContentDir:    utils.LookupEnv("CONTENT_DIR", DEFAULT_CONTENT_DIR),
		ContentFile:   utils.LookupEnv("MARKDOWN_CONTENT_FILE", DEFAULT_CONTENT_FILE),
	}
}
