package file

import (
	"image/jpeg"
	"image/png"
	"lazts/internal/models/types"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"lazts/internal/modules/log"

	"github.com/chai2010/webp"
)

func (h *handler) StaticFile(prefix string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		filePath := filepath.Join(prefix, r.URL.Path)
		ext := strings.ToLower(filepath.Ext(filePath))

		log.Fields("filePath", filePath, "url", r.URL.Path).D("Serving static file")

		switch ext {
		case ".css":
			h.minifyAndServeFile(w, filePath, types.CSS)
		case ".js":
			h.minifyAndServeFile(w, filePath, types.JavaScript)
		case ".svg":
			h.serveSVG(w, filePath)
		case ".jpeg", ".jpg":
			h.serveJPEG(w, filePath)
		case ".png":
			h.servePNG(w, filePath)
		case ".webp":
			h.serveWebP(w, filePath)
		default:
			http.ServeFile(w, r, filePath)
		}
	}
}

func (h *handler) minifyAndServeFile(w http.ResponseWriter, filePath, contentType string) {
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		http.Error(w, "Failed to read file", http.StatusInternalServerError)
		return
	}

	minifiedContent, err := h.minifier.Bytes(contentType, fileContent)
	if err != nil {
		http.Error(w, "Failed to minify content", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", contentType)
	w.Write(minifiedContent)
}

func (h *handler) serveJPEG(w http.ResponseWriter, filePath string) {
	img, err := h.watermaker.LoadImage(filePath)
	if err != nil {
		http.Error(w, "Failed to load image", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", types.JPEG)
	if err := jpeg.Encode(w, img, nil); err != nil {
		http.Error(w, "Failed to encode image", http.StatusInternalServerError)
	}
}

func (h *handler) servePNG(w http.ResponseWriter, filePath string) {
	img, err := h.watermaker.LoadImage(filePath)
	if err != nil {
		http.Error(w, "Failed to load image", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", types.PNG)
	if err := png.Encode(w, img); err != nil {
		http.Error(w, "Failed to encode image", http.StatusInternalServerError)
	}
}

func (h *handler) serveWebP(w http.ResponseWriter, filePath string) {
	img, err := h.watermaker.LoadImage(filePath)
	if err != nil {
		http.Error(w, "Failed to load image", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", types.WebP)
	if err := webp.Encode(w, img, nil); err != nil {
		http.Error(w, "Failed to encode image", http.StatusInternalServerError)
	}
}

func (h *handler) serveSVG(w http.ResponseWriter, filePath string) {
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		http.Error(w, "Failed to read file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", types.SVG)
	w.Write(fileContent)
}
