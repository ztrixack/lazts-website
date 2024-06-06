package api

import (
	"bytes"
	"lazts/internal/models"
	"lazts/internal/models/types"
	"net/http"
	"path/filepath"
	"strings"

	"lazts/internal/modules/log"

	"github.com/gorilla/mux"
)

func (h *handler) MemosGroupsContents(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	data := make(map[string]interface{})
	data["Menu"] = DEFAULT_MENU

	if params["content"] == "" {
		log.E("Contents is required")
		http.Error(w, "Contents is required", http.StatusBadRequest)
		return
	}

	data["Breadcrumbs"] = toMemoBreadcrumbs(params["group"])

	var buf bytes.Buffer
	if err := h.webber.RenderMarkdown(&buf, "memos-groups-contents", params["content"], data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	minifiedHTML, err := h.minifier.Bytes(types.HTML, buf.Bytes())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", types.HTML)
	w.Write(minifiedHTML)
}

func toMemoBreadcrumbs(tag string) []models.Tag {
	return []models.Tag{
		{
			Name: "Home",
			Link: "/",
		},
		{
			Name: "Memos",
			Link: "/memos",
		},
		{
			Name: strings.ToUpper(string(tag[0])) + tag[1:],
			Link: filepath.Join("/", "memos", tag),
		},
	}
}
