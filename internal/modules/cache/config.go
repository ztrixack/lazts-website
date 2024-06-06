package cache

import "lazts/internal/utils"

const (
	DEFAULT_TTL_SECONDS = 60
)

type config struct {
	TTLSeconds int
}

func parseConfig() *config {
	return &config{
		TTLSeconds: int(utils.LookupUIntEnv("CACHE_TTL_SECONDS", DEFAULT_TTL_SECONDS)),
	}
}
