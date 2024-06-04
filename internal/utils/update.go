package utils

import (
	"bytes"
	"fmt"
	"path/filepath"
	"regexp"
	"strings"
)

func UpdateImagePaths(markdownData []byte, prefixPath string) []byte {
	re := regexp.MustCompile(`!\[([^\]]*)\]\(([^)]+)\)`)

	replaceFunc := func(match []byte) []byte {
		parts := re.FindSubmatch(match)
		if len(parts) < 3 {
			return match
		}
		text := parts[1]
		url := parts[2]

		if bytes.HasPrefix(url, []byte("http://")) || bytes.HasPrefix(url, []byte("https://")) || bytes.HasPrefix(url, []byte("/")) {
			return match
		}

		newURL := filepath.Join(prefixPath, string(url))

		return []byte(fmt.Sprintf("![%s](%s)", text, newURL))
	}

	return re.ReplaceAllFunc(markdownData, replaceFunc)
}

func UpdateFeaturedImagePaths(path string, newPath string) string {
	if !strings.HasPrefix(path, "http://") && !strings.HasPrefix(path, "https://") {
		return filepath.Join(newPath, path)
	}

	return path
}
