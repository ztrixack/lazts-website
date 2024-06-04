package web

import (
	"io"
	"lazts/internal/modules/markdown"
	"lazts/internal/utils"
	"path/filepath"
	"text/template"

	"github.com/rs/zerolog/log"
)

type Servicer interface {
	RenderPage(w io.Writer, path string) error
	RenderMarkdown(w io.Writer, path string) error
}

type service struct {
	config    *config
	templates *template.Template
	markdown  markdown.Moduler
}

var _ Servicer = (*service)(nil)

func New(m markdown.Moduler) Servicer {
	c := parseConfig()
	t := parseTemplates(c.Dir)
	return &service{c, t, m}
}

func parseTemplates(path string) *template.Template {
	tmpl := template.New("")

	if _, err := utils.ParseAnyTemplates(tmpl, filepath.Join(path, "templates", "layouts/*.html")); err != nil {
		log.Fatal().Err(err).Msg("failed to parse layouts")
	}

	if _, err := utils.ParseAnyTemplates(tmpl, filepath.Join(path, "templates", "partials/*.html")); err != nil {
		log.Fatal().Err(err).Msg("failed to parse partials")
	}

	return tmpl
}
