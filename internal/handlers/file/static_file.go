package file

import "net/http"

func (h *handler) StaticFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path == "/manifest.json" {
		http.ServeFile(w, r, "web/static/root/manifest.json")
		return
	}

	if r.URL.Path == "/service-worker.js" {
		http.ServeFile(w, r, "web/static/root/service-worker.js")
		return
	}

	http.Error(w, "Resource not allowed", http.StatusNotFound)
}
