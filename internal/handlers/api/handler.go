package api

import (
	"lazts/internal/modules/http"
	"lazts/internal/services/book"
	"lazts/internal/services/memo"
	"lazts/internal/services/vacation"
	"lazts/internal/services/web"
)

type handler struct {
	webber     web.Servicer
	booker     book.Servicer
	vacationer vacation.Servicer
	memoizer   memo.Servicer
}

func New(m http.Moduler, ws web.Servicer, bs book.Servicer, vs vacation.Servicer, ms memo.Servicer) {
	h := &handler{ws, bs, vs, ms}
	h.setRouter(m)
}

func (h *handler) setRouter(m http.Moduler) {
	m.Get("/books", h.Books)
	m.Get("/about", h.About)
	m.Get("/", h.Home)
}
