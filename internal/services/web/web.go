package web

import (
	"io"
	"lazts/internal/modules/markdown"
	"lazts/internal/utils"
	"path/filepath"
	"text/template"
	"time"

	"lazts/internal/modules/log"
)

type Servicer interface {
	RenderPage(w io.Writer, path string, data map[string]interface{}) error
	RenderPartial(w io.Writer, path string, data map[string]interface{}) error
	RenderMarkdown(w io.Writer, path string, content string, data map[string]interface{}) error
	RenderVacationMarkdown(w io.Writer, path string, content string, data map[string]interface{}) error
}

type service struct {
	config    *config
	templates *template.Template
	markdown  markdown.Moduler
}

var _ Servicer = (*service)(nil)

func New(m markdown.Moduler) *service {
	c := parseConfig()
	t := parseTemplates(c.Dir)
	return &service{c, t, m}
}

func parseTemplates(path string) *template.Template {
	tmpl := template.New("")

	if _, err := utils.ParseAnyTemplates(tmpl, filepath.Join(path, "templates", "layouts/*.html")); err != nil {
		log.Err(err).C("failed to parse layouts")
	}

	if _, err := utils.ParseAnyTemplates(tmpl, filepath.Join(path, "templates", "partials/*.html")); err != nil {
		log.Err(err).C("failed to parse partials")
	}

	if _, err := utils.ParseAnyTemplates(tmpl, filepath.Join(path, "templates", "sections/**/*.html")); err != nil {
		log.Err(err).C("failed to parse sections")
	}

	return tmpl
}

func (m *service) injectData(data map[string]interface{}) map[string]interface{} {
	if data == nil {
		data = make(map[string]interface{})
	}

	data["XTitle"] = m.config.Title
	data["XExcerpt"] = m.config.Excerpt
	data["XYear"] = time.Now().Year()
	return data
}
