package api

import (
	"bytes"
	"lazts/internal/models/types"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

func (h *handler) MemosGroups(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	data := make(map[string]interface{})
	data["Menu"] = DEFAULT_MENU

	limit := uint(10)
	offset := uint(0)

	if page, err := strconv.Atoi(r.URL.Query().Get("page")); err == nil && page > 0 {
		offset = uint(page) * limit
	}

	if params["group"] == "" {
		log.Error().Msg("Group is required")
		http.Error(w, "Group is required", http.StatusBadRequest)
		return
	}

	var err error
	data["Memos"], err = h.memoizer.Get(offset, limit, params["group"])
	if err != nil {
		log.Error().Err(err).Msg("Failed to get memos")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var buf bytes.Buffer
	if err := h.webber.RenderPage(&buf, "memos-groups", data); err != nil {
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
