package cache

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseConfig(t *testing.T) {
	tests := []struct {
		name                string
		setup               func()
		teardown            func()
		expectedTTLSecondds int
	}{
		{
			name:                "environment variable set",
			expectedTTLSecondds: 10,
			setup: func() {
				os.Setenv("CACHE_TTL_SECONDS", "10")
			},
			teardown: func() {
				os.Unsetenv("CACHE_TTL_SECONDS")
			},
		},
		{
			name:                "environment variable not set",
			expectedTTLSecondds: DEFAULT_TTL_SECONDS,
			setup: func() {
				os.Unsetenv("CACHE_TTL_SECONDS")
			},
			teardown: func() {},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			defer tt.teardown()

			cfg := parseConfig()

			assert.Equal(t, tt.expectedTTLSecondds, cfg.TTLSeconds, "ttl secounds should be equal")
		})
	}
}
