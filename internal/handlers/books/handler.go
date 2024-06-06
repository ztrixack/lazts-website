package books

import (
	"lazts/internal/models/types"
	"lazts/internal/modules/http"
	"lazts/internal/services/book"
	"lazts/internal/services/web"

	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/html"
)

type handler struct {
	minifier *minify.M
	webber   web.Servicer
	booker   book.Servicer
}

func New(m http.Moduler, w web.Servicer, b book.Servicer) {
	mn := minify.New()
	mn.AddFunc(types.HTML, html.Minify)

	h := &handler{mn, w, b}
	h.setRouter(m)
}

func (h *handler) setRouter(m http.Moduler) {
	m.Get("/_books", h.List)
}
