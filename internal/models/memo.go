package models

import (
	"fmt"
	"lazts/internal/utils"
	"path/filepath"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
)

type MemoMetadata struct {
	Title         string   `json:"title"`
	Slug          string   `json:"slug"`
	Excerpt       string   `json:"excerpt"`
	FeaturedImage string   `json:"featured_image"`
	PublishedAt   string   `json:"published_at"`
	LastUpdatedAt string   `json:"last_updated_at"`
	Published     bool     `json:"published"`
	Tags          []string `json:"tags"`
	ReadTime      int      `json:"read_time"`
}

type Memo struct {
	Title               string
	Excerpt             string
	FeaturedImage       string
	Link                string
	Tags                []Tag
	ReadTime            int
	DateTimeISO         string
	DateTimeReadable    string
	LastUpdatedISO      string
	LastUpdatedReadable string
	DayMonth            string
	Year                string
}

type Tag struct {
	Name  string
	Link  string
	Count int
}

func (m MemoMetadata) ToMemo() Memo {
	publishedAt, err := time.Parse("2006-01-02", m.PublishedAt)
	if err != nil {
		publishedAt = time.Now()
	}
	lastUpdatedAt, err := time.Parse("2006-01-02", m.LastUpdatedAt)
	if err != nil {
		lastUpdatedAt = time.Now()
	}
	link := filepath.Join("/memos/main", m.Slug)
	if len(m.Tags) > 0 {
		link = filepath.Join("/memos", m.Tags[0], m.Slug)
	}

	return Memo{
		Title:               m.Title,
		Excerpt:             m.Excerpt,
		FeaturedImage:       utils.UpdateFeaturedImagePaths(filepath.Join("/static/contents/memos", m.Slug), m.FeaturedImage),
		Link:                link,
		Tags:                ToTags(m.Tags),
		ReadTime:            m.ReadTime,
		DateTimeISO:         publishedAt.Format(time.RFC3339),
		DateTimeReadable:    utils.ToYearMonthDay(publishedAt),
		LastUpdatedISO:      lastUpdatedAt.Format(time.RFC3339),
		LastUpdatedReadable: utils.ToYearMonthDay(lastUpdatedAt),
		DayMonth:            utils.ToDayMonth(publishedAt),
		Year:                fmt.Sprintf("%d", publishedAt.Year()),
	}
}

func ToTags(tags []string) []Tag {
	var t []Tag
	for _, tag := range tags {
		t = append(t, Tag{
			Name:  tag,
			Link:  filepath.Join("/memos", strings.ToLower(tag)),
			Count: 1,
		})
	}
	return t
}

func ToBreadcrumbs(tag string) []Tag {
	return []Tag{
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
			Link: filepath.Join("/memos", tag),
		},
	}
}

type MemoSort []Memo

func (m MemoSort) Len() int      { return len(m) }
func (m MemoSort) Swap(i, j int) { m[i], m[j] = m[j], m[i] }
func (m MemoSort) Less(i, j int) bool {
	// Convert DateTimeISO string to time.Time for comparison
	t1, err1 := time.Parse(time.RFC3339, m[i].DateTimeISO)
	t2, err2 := time.Parse(time.RFC3339, m[j].DateTimeISO)
	if err1 != nil || err2 != nil {
		log.Error().Err(err1).Err(err2).Msg("Error parsing date")
		return false
	}

	return t1.After(t2) // Sort in descending order
}

type TagSort []Tag

func (t TagSort) Len() int      { return len(t) }
func (t TagSort) Swap(i, j int) { t[i], t[j] = t[j], t[i] }
func (t TagSort) Less(i, j int) bool {
	return t[i].Name < t[j].Name
}
