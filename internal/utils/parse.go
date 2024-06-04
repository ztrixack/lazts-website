package utils

import (
	"path/filepath"
	"text/template"
)

func ParseAnyTemplates(tmpl *template.Template, path string) ([]string, error) {
	files, err := filepath.Glob(path)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		if _, err := tmpl.ParseFiles(file); err != nil {
			return nil, err
		}
	}

	return files, nil
}
