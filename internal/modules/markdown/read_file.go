package markdown

import (
	"lazts/internal/utils"
	"os"
	"path/filepath"
	"strings"
)

func (m *module) ReadFile(name, filename string) ([]byte, error) {
	KEY := "FILE-" + name
	if data, ok := m.cache[KEY].([]byte); data != nil && ok {
		return data, nil
	}

	content, ok := m.cache[KEY].([]byte)
	if ok && len(content) > 0 {
		return content, nil
	}

	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	parts := strings.Split(filename, "/")
	parts[0] = "/static"
	parts[len(parts)-1] = ""

	content = utils.UpdateImagePaths(content, filepath.Join(parts...))
	m.cache[KEY] = content

	return content, nil
}
