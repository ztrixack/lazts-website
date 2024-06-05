package models

import (
	"path/filepath"
	"time"

	"lazts/internal/utils"
)

type VacationMetadata struct {
	Title         string `json:"title"`
	Slug          string `json:"slug"`
	Excerpt       string `json:"excerpt"`
	Location      string `json:"location"`
	DateFrom      string `json:"date_from"`
	DateTo        string `json:"date_to"`
	FeaturedImage string `json:"featured_image"`
	PublishedAt   string `json:"published_at"`
	Published     bool   `json:"published"`
	LastUpdatedAt string `json:"last_updated_at"`
}

type Vacation struct {
	Title            string
	Excerpt          string
	Location         string
	DateTimeISO      string
	DateTimeReadable string
	FeaturedImage    string
	Link             string
}

func (v VacationMetadata) ToHTML() Vacation {
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
	}
}
