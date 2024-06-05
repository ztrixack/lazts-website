package books

import (
	"lazts/internal/modules/http"
	"lazts/internal/services/book"
	"lazts/internal/services/web"
)

type handler struct {
	webber web.Servicer
	book   book.Servicer
}

func New(m http.Moduler, w web.Servicer, b book.Servicer) {
	h := &handler{w, b}
	h.setRouter(m)
}

func (h *handler) setRouter(m http.Moduler) {
	m.Get("/_books", h.List)
}
