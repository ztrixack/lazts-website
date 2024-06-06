package book

import (
	"lazts/internal/modules/cache"
	"lazts/internal/utils"
	"os"
	"testing"
)

func setup(dir string) *service {
	os.Setenv("CONTENT_DIR", dir)

	r := New()
	r.cache = new(cache.Mock)

	return r
}

func teardown(t *testing.T, dir string) {
	utils.RemoveTestDir(t, dir)
	os.Unsetenv("CONTENT_DIR")
}
