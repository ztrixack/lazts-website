package file

import "net/http"

func (h *handler) Javascript(w http.ResponseWriter, r *http.Request) {
	fs := http.FileServer(http.Dir("web/static/js"))
	http.StripPrefix("/static/js", fs).ServeHTTP(w, r)
}
