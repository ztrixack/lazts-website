package books

import (
	"net/http"
)

func (h *handler) List(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("search")
	catalog := r.URL.Query().Get("catalog")
	status := r.URL.Query().Get("status")

	books, err := h.book.Get(search, catalog, status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := make(map[string]interface{})
	data["Books"] = books

	if err := h.webber.RenderPartial(w, "book-list", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
