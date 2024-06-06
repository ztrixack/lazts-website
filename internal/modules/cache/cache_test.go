package cache

import (
	"os"
	"strconv"
)

func setup(ttl int) *module {
	os.Setenv("CACHE_TTL_SECONDS", strconv.Itoa(ttl))
	return New()
}

func teardown() {
	os.Unsetenv("CACHE_TTL_SECONDS")
}
