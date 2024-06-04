package api

import (
	"net/http"

	"github.com/rs/zerolog/log"
)

func (h *handler) Home(w http.ResponseWriter, r *http.Request) {
	if err := h.webber.RenderPage(w, "home"); err != nil {
		log.Error().Err(err).Msg("failed to render page")
	}
}
