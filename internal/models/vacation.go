package models

import (
	"path/filepath"
	"strings"
	"time"

	"lazts/internal/utils"

	"lazts/internal/modules/log"
)

type VacationMetadata struct {
	Title         string   `json:"title"`
	Slug          string   `json:"slug"`
	Excerpt       string   `json:"excerpt"`
	Location      string   `json:"location"`
	DateFrom      string   `json:"date_from"`
	DateTo        string   `json:"date_to"`
	FeaturedImage string   `json:"featured_image"`
	PublishedAt   string   `json:"published_at"`
	Published     bool     `json:"published"`
	LastUpdatedAt string   `json:"last_updated_at"`
	Photos        []string `json:"photos"`
	Info          []string `json:"info"`
}

type Vacation struct {
	Title            string
	Excerpt          string
	Location         string
	DateTimeISO      string
	DateTimeReadable string
	FeaturedImage    string
	Link             string
	Photos           []string
	Info             []Option
}

func (v VacationMetadata) ToVacation() Vacation {
	from, err := time.Parse("2006-01-02", v.DateFrom)
	if err != nil {
		from = time.Now()
	}

	to, err := time.Parse("2006-01-02", v.DateTo)
	if err != nil {
		to = time.Now()
	}

	return Vacation{
		Title:            v.Title,
		Excerpt:          v.Excerpt,
		Location:         utils.ToFlagEmoji(v.Location),
		DateTimeISO:      from.Format(time.RFC3339),
		DateTimeReadable: utils.ToYearMonthDayRange(from, to),
		FeaturedImage:    utils.UpdateFeaturedImagePaths(filepath.Join("/static/contents/vacations", v.Slug), v.FeaturedImage),
		Link:             filepath.Join("/vacations", v.Slug),
		Photos:           v.Photos,
		Info:             getInfo(v.Info),
	}
}

type VacationSort []Vacation

func (v VacationSort) Len() int      { return len(v) }
func (v VacationSort) Swap(i, j int) { v[i], v[j] = v[j], v[i] }
func (v VacationSort) Less(i, j int) bool {
	// Convert DateTimeISO string to time.Time for comparison
	t1, err1 := time.Parse(time.RFC3339, v[i].DateTimeISO)
	t2, err2 := time.Parse(time.RFC3339, v[j].DateTimeISO)
	if err1 != nil || err2 != nil {
		log.Err(err1).Err(err2).E("Error parsing date")
		return false
	}

	return t1.After(t2) // Sort in descending order
}

func getInfo(info []string) []Option {
	result := make([]Option, 0)
	for _, i := range info {
		data := strings.Split(i, ";")
		if len(data) != 2 {
			continue
		}

		result = append(result, Option{
			Key:   data[0],
			Value: data[1],
		})
	}

	return result
}
