package web

import (
	"fmt"
	"io"
	"lazts/internal/models"
	"lazts/internal/utils"
	"os"
	"path/filepath"
	"strings"

	"lazts/internal/modules/log"
)

func (m *service) RenderMarkdown(w io.Writer, name string, content string, data map[string]interface{}) error {
	log.Fields("path", name).D("markdown rendering")

	tmpl, err := m.templates.Clone()
	if err != nil {
		log.Err(err).E("failed to clone templates")
		return ErrCloneTemplates
	}

	page, err := os.ReadFile(filepath.Join(m.config.Dir, "templates/pages", fmt.Sprintf("%s.html", name)))
	if err != nil {
		log.Err(err).E("failed to read file")
		return ErrNotFound
	}
	if _, err := tmpl.New("content").Parse(string(page)); err != nil {
		log.Err(err).E("failed to parse content")
		return ErrParseContent
	}

	htmlContent, err := m.markdown.LoadContent(strings.Split(name, "-")[0], content)
	if err != nil {
		log.Err(err).E("failed to convert markdown to html")
		return ErrConvertMarkdown
	}
	if _, err := tmpl.New("markdown").Parse(htmlContent); err != nil {
		log.Err(err).E("failed to parse content")
		return ErrParseContent
	}

	metadata, err := m.markdown.LoadMetadata(strings.Split(name, "-")[0], content)
	if err != nil {
		log.Err(err).E("failed to get metadata")
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
