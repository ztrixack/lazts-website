package web

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"lazts/internal/modules/log"
)

func (m *service) RenderPage(w io.Writer, name string, data map[string]interface{}) error {
	log.Fields("page", name).D("page rendering")

	tmpl, err := m.templates.Clone()
	if err != nil {
		log.Err(err).E("failed to clone templates")
		return ErrCloneTemplates
	}

	page, err := os.ReadFile(filepath.Join(m.config.Dir, "templates/pages", fmt.Sprintf("%s.html", name)))
	if err != nil {
		log.Err(err).E("failed to read file")
		return ErrNotFound
	}
	if _, err := tmpl.New("content").Parse(string(page)); err != nil {
		log.Err(err).E("failed to parse content")
		return ErrParseContent
	}

	return tmpl.ExecuteTemplate(w, "base", m.injectData(data))
}
