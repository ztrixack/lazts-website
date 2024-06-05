package markdown

import (
	"bytes"
	"lazts/internal/utils"
	"os"

	"github.com/rs/zerolog/log"
	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/parser"
)

func (m *module) ToMetadata(path string) (map[string]interface{}, error) {
	log.Debug().Str("path", path).Msg("converting to metadata")
	cpath := "metadata_" + path

	if data, ok := m.cache[cpath].(map[string]interface{}); data != nil && ok {
		log.Debug().Str("path", path).Msg("returning cached metadata for path")
		return data, nil
	}

	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	context := parser.NewContext()
	markdown := goldmark.New(goldmark.WithExtensions(meta.New()))

	buf := bytes.Buffer{}
	if err := markdown.Convert(content, &buf, parser.WithContext(context)); err != nil {
		return nil, err
	}

	metadata := meta.Get(context)
	metadata["read_time"] = utils.CalculateReadTime(content, m.config.WordPerMinute)

	m.cache[cpath] = metadata
	return metadata, nil
}
