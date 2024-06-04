package watermark

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	tests := []struct {
		name       string
		expectPath string
		expectSize int
	}{
		{"new config successfully", "./web/static/root/watermark.png", 48},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseConfig()
			assert.Equal(t, got.Path, tt.expectPath, "got %q, want %q", got.Path, tt.expectPath)
			assert.Equal(t, got.Size, tt.expectSize, "got %q, want %q", got.Size, tt.expectSize)
		})
	}
}
