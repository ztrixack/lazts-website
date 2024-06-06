package web

import (
	"lazts/internal/modules/markdown"
	"lazts/internal/utils"
	"os"
	"testing"
)

func setup(dir string, title string) (*service, *markdown.Mock) {
	os.Setenv("WEB_DIR", dir)
	os.Setenv("WEB_TITLE", title)

	mock := new(markdown.Mock)
	r := New(mock)

	return r, mock
}

func teardown(t *testing.T, dir string) {
	utils.RemoveTestDir(t, dir)
	os.Unsetenv("WEB_DIR")
	os.Unsetenv("WEB_TITLE")
}
