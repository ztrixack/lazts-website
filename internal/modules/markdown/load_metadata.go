package markdown

import (
	"bytes"
	"lazts/internal/utils"
	"path/filepath"

	"lazts/internal/modules/log"

	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/parser"
)

func (m *module) LoadMetadata(domain string, slug string) (map[string]interface{}, error) {
	KEY := m.genKey("METADATA", domain, slug)
	if content, found := m.cache.Get(KEY); found {
		return content.(map[string]interface{}), nil
	}

	filepath := filepath.Join(m.config.ContentDir, domain, slug, m.config.ContentFile)
	data, err := m.readFile(filepath)
	if err != nil {
		return nil, err
	}

	context := parser.NewContext()
	markdown := goldmark.New(goldmark.WithExtensions(meta.New()))

	var buf bytes.Buffer
	if err := markdown.Convert(data, &buf, parser.WithContext(context)); err != nil {
		return nil, err
	}

	metadata := meta.Get(context)
	if metadata != nil {
		metadata["read_time"] = utils.CalculateReadTime(data, m.config.WordPerMinute)
	} else {
		log.Fields("metadata", metadata).D("no metadata")
	}
	m.cache.Set(KEY, metadata)
	return metadata, nil
}
