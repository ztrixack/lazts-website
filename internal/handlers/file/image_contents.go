package file

import (
	"image/jpeg"
	"image/png"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/chai2010/webp"
)

func (h *handler) ImageContents(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	imagePath, _ := strings.CutPrefix(r.URL.Path, "/static")

	img, err := h.ws.LoadImage("./web" + imagePath)
	if err != nil {
		http.Error(w, "Failed to load image", http.StatusInternalServerError)
		return
	}

	switch strings.ToLower(filepath.Ext(imagePath)) {
	case ".jpeg", ".jpg":
		w.Header().Set("Content-Type", "image/jpeg")
		jpeg.Encode(w, img, &jpeg.Options{Quality: 80})
	case ".png":
		w.Header().Set("Content-Type", "image/png")
		png.Encode(w, img)
	case ".webp":
		w.Header().Set("Content-Type", "image/webp")
		webp.Encode(w, img, nil)
	default:
		http.Error(w, "Unsupported image format", http.StatusBadRequest)
		return
	}

}
