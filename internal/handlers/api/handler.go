package api

import (
	"lazts/internal/models/types"
	"lazts/internal/modules/http"
	"lazts/internal/services/book"
	"lazts/internal/services/memo"
	"lazts/internal/services/vacation"
	"lazts/internal/services/web"

	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/html"
)

type handler struct {
	minifier   *minify.M
	webber     web.Servicer
	booker     book.Servicer
	vacationer vacation.Servicer
	memoizer   memo.Servicer
}

var DEFAULT_MENU = []map[string]string{
	{"Label": "Home", "Path": "/"},
	{"Label": "Vacations", "Path": "/vacations"},
	{"Label": "Books", "Path": "/books"},
	{"Label": "Memos", "Path": "/memos"},
	{"Label": "About", "Path": "/about"},
}

func New(m http.Moduler, ws web.Servicer, bs book.Servicer, vs vacation.Servicer, ms memo.Servicer) {
	mn := minify.New()
	mn.AddFunc(types.HTML, html.Minify)

	h := &handler{mn, ws, bs, vs, ms}
	h.setRouter(m)
}

func (h *handler) setRouter(m http.Moduler) {
	m.Get("/about", h.About)
	m.Get("/books", h.Books)
	m.Get("/memos", h.Memos)
	m.Get("/memos/{group}", h.MemosGroups)
	m.Get("/vacations", h.Vacations)
	m.Get("/", h.Home)
}
