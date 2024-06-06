package api

import (
	"net/http"
	"strconv"

	"github.com/rs/zerolog/log"
)

func (h *handler) Memos(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	data["Menu"] = DEFAULT_MENU

	limit := uint(10)
	offset := uint(0)

	if page, err := strconv.Atoi(r.URL.Query().Get("page")); err == nil && page > 0 {
		offset = uint(page) * limit
	}

	var err error
	data["Memos"], err = h.memoizer.Get(offset, limit)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get memos")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data["Groups"], err = h.memoizer.GetTags()
	if err != nil {
		log.Error().Err(err).Msg("Failed to get tags")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := h.webber.RenderPage(w, "memos", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
