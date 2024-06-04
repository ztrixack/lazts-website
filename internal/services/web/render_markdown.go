package web

import (
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/rs/zerolog/log"
)

func (m *service) RenderMarkdown(w io.Writer, path string) error {
	log.Debug().Str("path", path).Msg("markdown rendering")

	tmpl, err := m.templates.Clone()
	if err != nil {
		log.Error().Err(err).Msg("failed to clone templates")
		return ErrCloneTemplates
	}

	pagename := strings.Split(strings.TrimPrefix(path, "/"), "/")[0]
	filename := pagename + "_content.html"

	page, err := os.ReadFile(filepath.Join(m.config.Dir, "templates/pages", filename))
	if err != nil {
		log.Error().Err(err).Msg("failed to read file")
		return ErrNotFound
	}
	if _, err := tmpl.New("content").Parse(string(page)); err != nil {
		log.Error().Err(err).Msg("failed to parse content")
		return ErrParseContent
	}

	htmlContent, err := m.markdown.ToHTML(filepath.Join(m.config.Dir, "contents", path, "page.md"))
	if err != nil {
		log.Error().Err(err).Msg("failed to convert markdown to html")
		return ErrConvertMarkdown
	}
	if _, err := tmpl.New("markdown").Parse(htmlContent); err != nil {
		log.Error().Err(err).Msg("failed to parse content")
		return ErrParseContent
	}

	return tmpl.ExecuteTemplate(w, "base", m.config)
}
