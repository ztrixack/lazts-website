package web

import (
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/rs/zerolog/log"
)

func (m *service) RenderPage(w io.Writer, path string) error {
	log.Debug().Str("path", path).Msg("page rendering")

	tmpl, err := m.templates.Clone()
	if err != nil {
		log.Error().Err(err).Msg("failed to clone templates")
		return ErrCloneTemplates
	}

	filename := strings.TrimPrefix(path, "/") + ".html"
	if path == "/" {
		filename = "home.html"
	}

	page, err := os.ReadFile(filepath.Join(m.config.Dir, "templates/pages", filename))
	if err != nil {
		log.Error().Err(err).Msg("failed to read file")
		return ErrNotFound
	}
	if _, err := tmpl.New("content").Parse(string(page)); err != nil {
		log.Error().Err(err).Msg("failed to parse content")
		return ErrParseContent
	}

	return tmpl.ExecuteTemplate(w, "base", m.config)
}
