package web

import (
	"io"

	"github.com/rs/zerolog/log"
)

func (m *service) RenderPartial(w io.Writer, name string, data map[string]interface{}) error {
	log.Debug().Str("partial", name).Msg("partial rendering")

	tmpl, err := m.templates.Clone()
	if err != nil {
		log.Error().Err(err).Msg("failed to clone templates")
		return ErrCloneTemplates
	}

	return tmpl.ExecuteTemplate(w, name, data)
}
