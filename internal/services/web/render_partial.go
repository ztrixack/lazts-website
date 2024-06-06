package web

import (
	"io"

	"lazts/internal/modules/log"
)

func (m *service) RenderPartial(w io.Writer, name string, data map[string]interface{}) error {
	log.Fields("partial", name).D("partial rendering")

	tmpl, err := m.templates.Clone()
	if err != nil {
		log.Err(err).E("failed to clone templates")
		return ErrCloneTemplates
	}

	return tmpl.ExecuteTemplate(w, name, data)
}
