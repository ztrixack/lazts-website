package markdown

import (
	_ "embed"
)

//go:embed mermaid.min.js
var mermaidJSSource string

type Moduler interface {
	ToHTML(filepath string) (string, error)
	ToMetadata(filepath string) (map[string]interface{}, error)
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
