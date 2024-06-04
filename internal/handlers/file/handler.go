package file

import (
	"lazts/internal/modules/http"
	"lazts/internal/services/watermark"
)

type handler struct {
	ws watermark.Servicer
}

func New(m http.Moduler, ws watermark.Servicer) {
	h := &handler{ws}
	h.setRouter(m)
}

func (h *handler) setRouter(m http.Moduler) {
	m.Get("/static/icons/", h.Icons)
	m.Get("/static/images/", h.Images)
	m.Get("/static/contents/", h.ImageContents)
	m.Get("/static/js/", h.Javascript)
	m.Get("/static/css/", h.CSS)
	m.Get("/manifest.json", h.StaticFile)
	m.Get("/service-worker.js", h.StaticFile)
}
