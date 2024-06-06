package api

import (
	"net/http"
)

func (h *handler) Vacations(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	data["Menu"] = DEFAULT_MENU

	var err error
	data["Vacations"], err = h.vacationer.Get("")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := h.webber.RenderPage(w, "vacations", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
