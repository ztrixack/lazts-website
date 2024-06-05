package api

import (
	"lazts/internal/modules/http"
	"lazts/internal/services/book"
	"lazts/internal/services/vacation"
	"lazts/internal/services/web"
)

type handler struct {
	webber     web.Servicer
	booker     book.Servicer
	vacationer vacation.Servicer
}

func New(m http.Moduler, ws web.Servicer, bs book.Servicer, vs vacation.Servicer) {
	h := &handler{ws, bs, vs}
	h.setRouter(m)
}

func (h *handler) setRouter(m http.Moduler) {
	m.Get("/", h.Home)
	m.Get("/about", h.About)
}
