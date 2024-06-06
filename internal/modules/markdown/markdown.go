package markdown

import (
	_ "embed"
	"lazts/internal/modules/cache"
	"lazts/internal/utils"
	"os"
	"path/filepath"
	"strings"
)

//go:embed mermaid.min.js
var mermaidJSSource string

type Moduler interface {
	LoadContent(domain string, slug string) (string, error)
	LoadMetadata(domain string, slug string) (map[string]interface{}, error)
}

type module struct {
	config *config
	cache  cache.Moduler
}

var _ Moduler = (*module)(nil)

func New() *module {
	return &module{
		config: parseConfig(),
		cache:  cache.New(),
	}
}

func (m *module) readFile(fullpath string) ([]byte, error) {
	KEY := m.genKey("FILE", fullpath, "0")
	if content, found := m.cache.Get(KEY); found {
		return content.([]byte), nil
	}

	content, err := os.ReadFile(fullpath)
	if err != nil {
		return nil, err
	}

	parts := strings.Split(fullpath, "/")
	parts[0] = "/static"
	parts[len(parts)-1] = ""

	content = utils.UpdateImagePaths(content, filepath.Join(parts...))
	m.cache.Set(KEY, content)

	return content, nil
}

func (m *module) genKey(fn, domain, slug string) string {
	return fn + "-" + domain + "-" + slug
}
