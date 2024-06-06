package api

import (
	"bytes"
	"lazts/internal/models"
	"lazts/internal/models/types"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

func (h *handler) VacationsContents(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	data := make(map[string]interface{})
	data["Menu"] = DEFAULT_MENU

	if params["content"] == "" {
		log.Error().Msg("Contents is required")
		http.Error(w, "Contents is required", http.StatusBadRequest)
		return
	}

	data["Breadcrumbs"] = toVacationBreadcrumbs()

	var buf bytes.Buffer
	if err := h.webber.RenderMarkdown(&buf, "vacations-contents", params["content"], data); err != nil {
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

func toVacationBreadcrumbs() []models.Tag {
	return []models.Tag{
		{
			Name: "Home",
			Link: "/",
		},
		{
			Name: "Vacations",
			Link: "/vacations",
		},
	}
}