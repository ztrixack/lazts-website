package web

import (
	"fmt"
	"io"
	"lazts/internal/models"
	"lazts/internal/utils"
	"os"
	"path/filepath"
	"strings"

	"github.com/rs/zerolog/log"
)

func (m *service) RenderMarkdown(w io.Writer, name string, content string, data map[string]interface{}) error {
	log.Debug().Str("path", name).Msg("markdown rendering")

	tmpl, err := m.templates.Clone()
	if err != nil {
		log.Error().Err(err).Msg("failed to clone templates")
		return ErrCloneTemplates
	}

	page, err := os.ReadFile(filepath.Join(m.config.Dir, "templates/pages", fmt.Sprintf("%s.html", name)))
	if err != nil {
		log.Error().Err(err).Msg("failed to read file")
		return ErrNotFound
	}
	if _, err := tmpl.New("content").Parse(string(page)); err != nil {
		log.Error().Err(err).Msg("failed to parse content")
		return ErrParseContent
	}

	filepath := filepath.Join(m.config.Dir, "contents", strings.Split(name, "-")[0], content, "index.md")
	filedata, err := m.markdown.ReadFile(content, filepath)
	if err != nil {
		log.Error().Err(err).Msg("failed to read file")
		return ErrNotFound
	}

	htmlContent, err := m.markdown.ToHTML(content, filedata)
	if err != nil {
		log.Error().Err(err).Msg("failed to convert markdown to html")
		return ErrConvertMarkdown
	}
	if _, err := tmpl.New("markdown").Parse(htmlContent); err != nil {
		log.Error().Err(err).Msg("failed to parse content")
		return ErrParseContent
	}

	metadata, err := m.markdown.ToMetadata(content, filedata)
	if err != nil {
		log.Error().Err(err).Msg("failed to get metadata")
		return err
	}
	var metamemo models.MemoMetadata
	if err := utils.ToStruct(metadata, &metamemo); err != nil {
		return err
	}
	memo := metamemo.ToMemo()

	data["Title"] = memo.Title
	data["Published"] = memo.DateTimeReadable
	data["PublishedISO"] = memo.DateTimeISO
	data["LastUpdated"] = memo.LastUpdatedReadable
	data["LastUpdatedISO"] = memo.LastUpdatedISO
	data["ReadTime"] = memo.ReadTime
	data["Tags"] = memo.Tags

	return tmpl.ExecuteTemplate(w, "base", m.injectData(data))
}
