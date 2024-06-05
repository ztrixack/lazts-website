package models

import (
	"fmt"
	"lazts/internal/utils"
	"path/filepath"
	"strings"
	"time"
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
	Title            string
	Excerpt          string
	FeaturedImage    string
	Link             string
	Tags             []Tag
	ReadTime         int
	DateTimeISO      string
	DateTimeReadable string
	DayMonth         string
	Year             string
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

	return Memo{

		Title:            m.Title,
		Excerpt:          m.Excerpt,
		FeaturedImage:    utils.UpdateFeaturedImagePaths(filepath.Join("/static/contents/memos", m.Slug), m.FeaturedImage),
		Link:             filepath.Join("/memos", m.Tags[0], m.Slug),
		Tags:             ToTags(m.Tags),
		ReadTime:         m.ReadTime,
		DateTimeISO:      publishedAt.Format(time.RFC3339),
		DateTimeReadable: utils.ToYearMonthDay(publishedAt),
		DayMonth:         utils.ToDayMonth(publishedAt),
		Year:             fmt.Sprintf("%d", publishedAt.Year()),
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
