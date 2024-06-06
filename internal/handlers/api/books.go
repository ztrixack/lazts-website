package api

import (
	"net/http"
)

func (h *handler) Books(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	data["Menu"] = DEFAULT_MENU

	search := r.URL.Query().Get("search")
	catalog := r.URL.Query().Get("catalog")
	status := r.URL.Query().Get("status")

	var err error
	data["Books"], err = h.booker.Get(search, catalog, status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data["Catalogs"], err = h.booker.GetCatalogs()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data["Size"] = h.booker.GetSize()
	data["CurrentCatalog"] = catalog
	data["CurrentStatus"] = status

	if err := h.webber.RenderPage(w, "books", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
