package watermark

import (
	"lazts/internal/modules/cache"
	"lazts/internal/modules/imaging"
	"lazts/internal/utils"
	"os"
	"testing"
)

func setup(t *testing.T, dir string) (*service, *imaging.Mock) {
	os.Setenv("WEB_DIR", dir)
	os.Setenv("WATERMARK_PATH", dir+"/watermark.png")
	os.Setenv("WATERMARK_SIZE", "32")

	utils.CreateTestFile(t, dir, "original.png", "")
	utils.CreateTestFile(t, dir, "watermark.png", "")

	mock := new(imaging.Mock)
	r := New(mock)
	r.cache = new(cache.Mock)

	return r, mock
}

func teardown(t *testing.T, dir string) {
	utils.RemoveTestDir(t, dir)
	os.Unsetenv("WEB_DIR")
	os.Unsetenv("WATERMARK_PATH")
	os.Unsetenv("WATERMARK_SIZE")
}
