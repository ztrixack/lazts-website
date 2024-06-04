package http

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	tests := []struct {
		name         string
		env          map[string]string
		expectedHost string
		expectedPort string
	}{
		{
			name:         "Given HTTP_HOST set to '', HTTP_PORT set to 8082, expect empty host and port 8082",
			env:          map[string]string{"HTTP_HOST": "", "HTTP_PORT": "8082"},
			expectedHost: "",
			expectedPort: "8082",
		},
		{
			name:         "Given HTTP_PORT set to 8082, expect localhost and port 8082",
			env:          map[string]string{"HTTP_PORT": "8082"},
			expectedHost: "localhost",
			expectedPort: "8082",
		},
		{
			name:         "Given no set, expect default empty host and port 8080",
			env:          map[string]string{},
			expectedHost: "localhost",
			expectedPort: "8080",
		},
		{
			name:         "Given HTTP_PORT set to non-numeric, expect default port 8080",
			env:          map[string]string{"HTTP_PORT": "abc"},
			expectedHost: "localhost",
			expectedPort: "8080",
		},
		{
			name:         "Given HTTP_HOST set to any, expect empty host",
			env:          map[string]string{"HTTP_HOST": "new-host"},
			expectedHost: "new-host",
			expectedPort: "8080",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d: %s", i+1, tt.name), func(t *testing.T) {
			for key, value := range tt.env {
				os.Setenv(key, value)
				defer os.Unsetenv(key)
			}

			c := parseConfig()

			assert.Equal(t, tt.expectedHost, c.Host, "The host configuration does not match the expected value.")
			assert.Equal(t, tt.expectedPort, c.Port, "The port configuration does not match the expected value.")
		})
	}
}
