package api

import (
	"bytes"
	"lazts/internal/models/types"
	"net/http"
)

func (h *handler) About(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	data["Menu"] = DEFAULT_MENU
	data["Vacations"] = h.vacationer.GetSize()
	data["Books"] = h.booker.GetSize()
	data["Memos"] = h.memoizer.GetSize()

	var buf bytes.Buffer
	if err := h.webber.RenderPage(&buf, "about", data); err != nil {
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
