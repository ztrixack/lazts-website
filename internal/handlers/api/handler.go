package api

import (
	"lazts/internal/modules/http"
	"lazts/internal/services/web"
)

type handler struct {
	webber web.Servicer
}

func New(m http.Moduler, ws web.Servicer) {
	h := &handler{ws}
	h.setRouter(m)
}

func (h *handler) setRouter(m http.Moduler) {
	// page
	m.Get("/", h.Home)
}
