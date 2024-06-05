package api

import (
	"lazts/internal/utils"
	"net/http"

	"github.com/rs/zerolog/log"
)

func (h *handler) Home(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	data["Menu"] = DEFAULT_MENU
	data["Since"] = "Since 1991"
	data["Blackhole"] = utils.RandomizeBlackholes(1000)
	data["Cloud"] = utils.RandomizeClouds(100)

	var err error
	data["BookStats"], err = h.booker.GetStats()
	if err != nil {
		log.Error().Err(err).Msg("failed to get book stats")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data["BookShelf"], err = h.booker.GetShelf(2, 3, 4)
	if err != nil {
		log.Error().Err(err).Msg("failed to get random 2 books")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data["Vacations"], err = h.vacationer.Get("japan")
	if err != nil {
		log.Error().Err(err).Msg("failed to get vacations")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data["Memos"], err = h.memoizer.Get(0, 10)
	if err != nil {
		log.Error().Err(err).Msg("failed to get memos")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := h.webber.RenderPage(w, "home", data); err != nil {
		log.Error().Err(err).Msg("failed to render page")
	}
}

var DEFAULT_MENU = []map[string]string{
	{"Label": "Home", "Path": "/"},
	{"Label": "Vacations", "Path": "/vacations"},
	{"Label": "Books", "Path": "/books"},
	{"Label": "Memos", "Path": "/memos"},
	{"Label": "About", "Path": "/about"},
}
