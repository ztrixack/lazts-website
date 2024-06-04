package api

import (
	"math"
	"math/rand/v2"
	"net/http"

	"github.com/rs/zerolog/log"
)

func (h *handler) Home(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	data["Menu"] = DEFAULT_MENU
	data["Since"] = "Since 1991"
	data["Blackhole"] = randomizeBlackholes(1000)

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

type Blackhole struct {
	Size    int
	Rotate  int
	Opacity int
	Width   int
}

func randomizeBlackholes(count int) []Blackhole {
	var blackholes []Blackhole
	for i := 0; i < count; i++ {
		size := rand.IntN(360) + 180
		rotate := rand.IntN(360)
		opacity := int(math.Max(110-float64(size*100/450), 5))
		width := (size - 90) / 6

		blackholes = append(blackholes, Blackhole{Size: size, Rotate: rotate, Opacity: opacity, Width: width})
	}
	return blackholes
}
