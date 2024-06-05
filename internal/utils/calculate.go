package utils

import (
	"regexp"
	"strings"
)

func CalculateReadTime(markdown []byte, wpm int) int {
	text := removeMarkdownFormatting(string(markdown))
	words := strings.Fields(text)
	wordCount := len(words)
	readTime := wordCount / wpm
	if wordCount%wpm != 0 {
		readTime++
	}

	return readTime
}

func removeMarkdownFormatting(text string) string {
	urlRegex := regexp.MustCompile(`\[(.*?)\]\(http.*?\)`)
	text = urlRegex.ReplaceAllString(text, "$1")
	symbols := []string{"#", "*", "_", "`"}
	for _, symbol := range symbols {
		text = strings.ReplaceAll(text, symbol, "")
	}

	return text
}
