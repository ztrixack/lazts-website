package web

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/rs/zerolog/log"
)

func (m *service) RenderPage(w io.Writer, name string, data map[string]interface{}) error {
	log.Debug().Str("page", name).Msg("page rendering")

	tmpl, err := m.templates.Clone()
	if err != nil {
		log.Error().Err(err).Msg("failed to clone templates")
		return ErrCloneTemplates
	}

	page, err := os.ReadFile(filepath.Join(m.config.Dir, "templates/pages", fmt.Sprintf("%s.html", name)))
	if err != nil {
		log.Error().Err(err).Msg("failed to read file")
		return ErrNotFound
	}
	if _, err := tmpl.New("content").Parse(string(page)); err != nil {
		log.Error().Err(err).Msg("failed to parse content")
		return ErrParseContent
	}

	log.Debug().Str("page", name).Msg("rendering page")

	return tmpl.ExecuteTemplate(w, "base", m.injectData(data))
}
