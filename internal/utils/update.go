package utils

import (
	"bytes"
	"fmt"
	"net/url"
	"path/filepath"
	"regexp"
	"strings"
)

func UpdateHTMLImagePaths(markdownData []byte, prefixPath string) []byte {
	re := regexp.MustCompile(`<img[^>]+src=["']([^"']+)["']`)

	replaceFunc := func(match []byte) []byte {
		parts := re.FindSubmatch(match)
		if len(parts) < 2 {
			return match
		}
		url := parts[1]

		if bytes.HasPrefix(url, []byte("http://")) || bytes.HasPrefix(url, []byte("https://")) || bytes.HasPrefix(url, []byte("/")) {
			return match
		}

		newURL := filepath.Join(prefixPath, string(url))

		return bytes.Replace(match, url, []byte(newURL), 1)
	}

	return re.ReplaceAllFunc(markdownData, replaceFunc)
}

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

func UpdateFeaturedImagePaths(base string, path string) string {
	if !strings.HasPrefix(path, "http://") && !strings.HasPrefix(path, "https://") {
		urlpath, err := url.JoinPath(base, path)
		if err != nil {
			return path
		}

		return urlpath
	}

	return path
}
