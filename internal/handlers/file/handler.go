package file

import (
	"lazts/internal/models/types"
	"lazts/internal/modules/http"
	"lazts/internal/services/watermark"

	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/js"
)

type handler struct {
	minifier   *minify.M
	watermaker watermark.Servicer
}

func New(m http.Moduler, ws watermark.Servicer) {
	mn := minify.New()
	mn.AddFunc(types.CSS, css.Minify)
	mn.AddFunc(types.JavaScript, js.Minify)

	h := &handler{mn, ws}
	h.setRouter(m)
}

func (h *handler) setRouter(m http.Moduler) {
	m.StaticFileServer("/static/icons/", h.IconFile("./web/static/icons"))
	m.StaticFileServer("/static/svg/", h.StaticFile("./web/static/svg"))
	m.StaticFileServer("/static/css/", h.StaticFile("./web/static/css"))
	m.StaticFileServer("/static/js/", h.StaticFile("./web/static/js"))
	m.StaticFileServer("/static/images/", h.StaticFile("./static/images"))
	m.StaticFileServer("/static/contents/", h.StaticFile("./contents"))
	m.Get("/manifest.json", h.StaticFile("./web/static/root"))
	m.Get("/service-worker.js", h.StaticFile("./webstatic/root"))
}
