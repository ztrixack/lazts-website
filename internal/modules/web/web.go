package web

import (
	"io"
	"lazts/internal/utils"
	"path/filepath"
	"text/template"

	"github.com/rs/zerolog/log"
)

type Webber interface {
	RenderPage(w io.Writer, path string) error
	RenderMarkdown(w io.Writer, path string) error
}

type webber struct {
	config    *config
	templates *template.Template
}

func New() Webber {
	c := parseConfig()
	t := parseTemplates(c.Dir)
	return &webber{c, t}
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
