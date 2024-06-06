package markdown

import (
	"bytes"
	"lazts/internal/utils"

	"github.com/rs/zerolog/log"
	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/parser"
)

func (m *module) ToMetadata(name string, data []byte) (map[string]interface{}, error) {
	KEY := "METADATA-" + name
	if data, ok := m.cache[KEY].(map[string]interface{}); data != nil && ok {
		return data, nil
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
		log.Debug().Interface("metadata", metadata).Msg("no metadata")
	}

	m.cache[KEY] = metadata
	return metadata, nil
}
