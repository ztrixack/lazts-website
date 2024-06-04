package web

import (
	"lazts/internal/utils"
)

const (
	DEFAULT_DIR     = "./web"
	DEFAULT_TITLE   = "lazts"
	DEFAULT_EXCERPT = "no excerpt"
)

type config struct {
	Dir     string
	Title   string
	Excerpt string
}

func parseConfig() *config {
	return &config{
		Dir:     utils.LookupEnv("WEB_DIR", DEFAULT_DIR),
		Title:   utils.LookupEnv("WEB_TITLE", DEFAULT_TITLE),
		Excerpt: utils.LookupEnv("WEB_EXCERPT", DEFAULT_EXCERPT),
	}
}
