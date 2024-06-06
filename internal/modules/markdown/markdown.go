package markdown

import (
	_ "embed"
)

//go:embed mermaid.min.js
var mermaidJSSource string

type Moduler interface {
	ReadFile(name string, filepath string) ([]byte, error)
	ToHTML(name string, data []byte) (string, error)
	ToMetadata(name string, data []byte) (map[string]interface{}, error)
}

type module struct {
	config *config
	cache  map[string]interface{}
}

var _ Moduler = (*module)(nil)

func New() *module {
	return &module{
		config: parseConfig(),
		cache:  make(map[string]interface{}),
	}
}
