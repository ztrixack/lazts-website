package books

import (
	"bytes"
	"lazts/internal/models/types"
	"net/http"
)

func (h *handler) List(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("search")
	catalog := r.URL.Query().Get("catalog")
	status := r.URL.Query().Get("status")

	books, err := h.booker.Get(search, catalog, status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := make(map[string]interface{})
	data["Books"] = books

	var buf bytes.Buffer
	if err := h.webber.RenderPartial(&buf, "book-list", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	minifiedHTML, err := h.minifier.Bytes(types.HTML, buf.Bytes())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", types.HTML)
	w.Write(minifiedHTML)
}
