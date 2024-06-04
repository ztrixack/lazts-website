package file

import "net/http"

func (h *handler) Icons(w http.ResponseWriter, r *http.Request) {
	fs := http.FileServer(http.Dir("web/static/icons"))
	http.StripPrefix("/static/icons", fs).ServeHTTP(w, r)
}
