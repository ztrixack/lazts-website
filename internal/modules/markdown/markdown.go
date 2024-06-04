package markdown

import (
	_ "embed"
)

//go:embed mermaid.min.js
var mermaidJSSource string

type Moduler interface {
	ToHTML(filepath string) (string, error)
}

type module struct {
	config *config
	cache  map[string]string
}

var _ Moduler = (*module)(nil)

func New() Moduler {
	return &module{
		config: parseConfig(),
		cache:  make(map[string]string),
	}
}
