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

	if err := h.webber.RenderPage(w, "home", data); err != nil {
		log.Error().Err(err).Msg("failed to render page")
	}
}

var DEFAULT_MENU = []map[string]string{
	{"Label": "Home", "Path": "/"},
	{"Label": "Vacations", "Path": "/vacations"},
	{"Label": "Books", "Path": "/books"},
	{"Label": "Notes", "Path": "/notes"},
	{"Label": "About", "Path": "/about"},
}
