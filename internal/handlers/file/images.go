package file

import (
	"image/jpeg"
	"net/http"
	"strings"
)

func (h *handler) Images(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if !strings.HasSuffix(r.URL.Path, ".jpeg") {
		fs := http.FileServer(http.Dir("web/static/images"))
		http.StripPrefix("/static/images", fs).ServeHTTP(w, r)
		return
	}

	img, err := h.ws.LoadImage(r.URL.Path)
	if err != nil {
		http.Error(w, "Failed to load image", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/jpeg")

	jpeg.Encode(w, img, nil)
}
