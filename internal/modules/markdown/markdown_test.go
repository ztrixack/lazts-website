package markdown

import (
	"lazts/internal/modules/cache"
	"lazts/internal/utils"
	"os"
	"testing"
)

func setup(t *testing.T) *module {
	os.Setenv("CONTENT_DIR", "./test_data")
	os.Setenv("MARKDOWN_CONTENT_FILE", "test.md")

	m := New()
	m.cache = new(cache.Mock)

	utils.CreateTestFile(t, "test_data/domain/slug", "test.md", `---
meta: test
---
some content`)

	return m
}

func teardown(t *testing.T) {
	utils.RemoveTestDir(t, "test_data")
	os.Unsetenv("CONTENT_DIR")
	os.Unsetenv("MARKDOWN_CONTENT_FILE")
}
