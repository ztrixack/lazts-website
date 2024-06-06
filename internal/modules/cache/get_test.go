package cache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCache_Get(t *testing.T) {
	m := setup(1)
	defer teardown()

	tests := []struct {
		name   string
		key    string
		value  interface{}
		found  bool
		sleep  time.Duration
		result interface{}
	}{
		{"Success", "key1", "value1", true, 0, "value1"},
		{"Out of time", "key1", "value1", false, 2 * time.Second, nil},
	}

	for _, tt := range tests {
		m.Set(tt.key, tt.value)
		if tt.sleep > 0 {
			time.Sleep(tt.sleep)
		}
		value, found := m.Get(tt.key)
		assert.Equal(t, tt.found, found)
		assert.Equal(t, tt.result, value)
	}
}
