package api

import (
	"lazts/internal/modules/http"
	"lazts/internal/services/book"
	"lazts/internal/services/web"
)

type handler struct {
	webber web.Servicer
	booker book.Servicer
}

func New(m http.Moduler, ws web.Servicer, bs book.Servicer) {
	h := &handler{ws, bs}
	h.setRouter(m)
}

func (h *handler) setRouter(m http.Moduler) {
	m.Get("/", h.Home)
	m.Get("/about", h.About)
}
