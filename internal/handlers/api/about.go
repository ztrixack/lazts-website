package api

import "net/http"

func (h *handler) About(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	data["Menu"] = DEFAULT_MENU

	if err := h.webber.RenderPage(w, "about", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
