package file

import "net/http"

func (h *handler) CSS(w http.ResponseWriter, r *http.Request) {
	fs := http.FileServer(http.Dir("web/static/css"))
	http.StripPrefix("/static/css", fs).ServeHTTP(w, r)
}
