package markdown

import (
	"bytes"
	"os"
	"strings"

	katex "github.com/FurqanSoftware/goldmark-katex"
	chromahtml "github.com/alecthomas/chroma/v2/formatters/html"
	figure "github.com/mangoumbrella/goldmark-figure"
	"github.com/rs/zerolog/log"
	"github.com/yuin/goldmark"
	emoji "github.com/yuin/goldmark-emoji"
	_ "github.com/yuin/goldmark-emoji/definition"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	"go.abhg.dev/goldmark/mermaid"
	"go.abhg.dev/goldmark/mermaid/mermaidcdp"
	"go.abhg.dev/goldmark/toc"
)

func (m *module) ToHTML(path string) (string, error) {
	log.Debug().Str("path", path).Msg("converting to HTML")

	if m.cache[path] != "" {
		log.Debug().Str("path", path).Msg("returning cached HTML for path")
		return m.cache[path], nil
	}

	content, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	context := parser.NewContext()
	compiler, err := mermaidcdp.New(&mermaidcdp.Config{JSSource: mermaidJSSource})
	if err != nil {
		return "", err
	}
	defer compiler.Close()

	markdown := goldmark.New(
		goldmark.WithExtensions(
			extension.GFM, emoji.Emoji, figure.Figure, extension.Table,
			meta.New(meta.WithStoresInDocument()),
			highlighting.NewHighlighting(
				highlighting.WithStyle(m.config.Theme),
				highlighting.WithFormatOptions(
					chromahtml.WithLineNumbers(true),
				),
			),
			&katex.Extender{}, &toc.Extender{ListID: "toc-list", TitleDepth: 2},
			&mermaid.Extender{Compiler: compiler}),
		goldmark.WithParserOptions(parser.WithAutoHeadingID()),
		goldmark.WithRendererOptions(html.WithHardWraps(), html.WithXHTML(), html.WithUnsafe()),
	)

	buf := bytes.Buffer{}
	if err := markdown.Convert(content, &buf, parser.WithContext(context)); err != nil {
		return "", err
	}

	m.cache[path] = strings.TrimSpace(buf.String())
	return m.cache[path], nil
}
