package cache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCache_Set(t *testing.T) {
	m := setup(1)
	defer teardown()

	tests := []struct {
		key   string
		value interface{}
	}{
		{"key1", "value1"},
		{"key2", "value2"},
	}

	for _, tc := range tests {
		m.Set(tc.key, tc.value)
		m.mu.RLock()
		item, found := m.data[tc.key]
		m.mu.RUnlock()
		assert.True(t, found)
		assert.Equal(t, tc.value, item.Value)
	}
}
