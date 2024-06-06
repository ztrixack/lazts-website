package cache

import "time"

func (m *module) Get(key string) (interface{}, bool) {
	m.mu.RLock()
	item, found := m.data[key]
	m.mu.RUnlock()
	if !found || time.Now().UnixNano() > item.Expiration {
		return nil, false
	}
	return item.Value, true
}
