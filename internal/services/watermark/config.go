package watermark

import (
	"lazts/internal/utils"
)

const (
	DEFAULT_DIR  = "./web"
	DEFAULT_PATH = "./web/static/root/watermark.png"
	DEFAULT_SIZE = 48
)

type config struct {
	Dir  string
	Path string
	Size int
}

func parseConfig() *config {
	return &config{
		Dir:  utils.LookupEnv("WEB_DIR", DEFAULT_DIR),
		Path: utils.LookupEnv("WATERMARK_PATH", DEFAULT_PATH),
		Size: int(utils.LookupUIntEnv("WATERMARK_SIZE", DEFAULT_SIZE)),
	}
}
