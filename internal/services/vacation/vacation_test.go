package vacation

import (
	"lazts/internal/modules/cache"
	"lazts/internal/modules/markdown"
	"lazts/internal/utils"
	"os"
	"testing"
)

func setup(dir string) (*service, *markdown.Mock) {
	os.Setenv("CONTENT_DIR", dir)

	mock := new(markdown.Mock)
	r := New(mock)
	r.cache = new(cache.Mock)

	return r, mock
}

func teardown(t *testing.T, dir string) {
	utils.RemoveTestDir(t, dir)
	os.Unsetenv("CONTENT_DIR")
}
