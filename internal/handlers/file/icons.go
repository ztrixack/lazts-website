package file

import (
	"net/http"
	"path/filepath"
)

func (h *handler) IconFile(prefix string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		filePath := filepath.Join(prefix, r.URL.Path)
		http.ServeFile(w, r, filePath)
	}
}
