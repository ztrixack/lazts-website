package utils

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateHTMLImagePaths(t *testing.T) {
	tests := []struct {
		input    string
		prefix   string
		expected string
	}{
		{
			input:    `<div class="gallery"><img src="iceland/bollard-01.png" /><img src="iceland/bollard-02.png" /></div>`,
			prefix:   "/static/contents/memos/20240616-geoguessr-on-the-road",
			expected: `<div class="gallery"><img src="/static/contents/memos/20240616-geoguessr-on-the-road/iceland/bollard-01.png" /><img src="/static/contents/memos/20240616-geoguessr-on-the-road/iceland/bollard-02.png" /></div>`,
		},
		{
			input:    `<img src="image.png" />`,
			prefix:   "/static/images",
			expected: `<img src="/static/images/image.png" />`,
		},
		{
			input:    `<img src="/already/prefixed/image.png" />`,
			prefix:   "/static/images",
			expected: `<img src="/already/prefixed/image.png" />`,
		},
		{
			input:    `<img src="http://example.com/image.png" />`,
			prefix:   "/static/images",
			expected: `<img src="http://example.com/image.png" />`,
		},
	}

	for _, tt := range tests {
		result := UpdateHTMLImagePaths([]byte(tt.input), tt.prefix)
		if !bytes.Equal(result, []byte(tt.expected)) {
			assert.Equal(t, tt.expected, string(result), "output should match expected")
		}
	}
}

func TestUpdateImagePaths(t *testing.T) {
	tests := []struct {
		name     string
		markdown string
		expected string
	}{
		{
			name:     "Normal Image Paths",
			markdown: `This is a test image: ![test image](images/test.png) and ![another image](images/2021/img.jpg)`,
			expected: `This is a test image: ![test image](/new/path/images/test.png) and ![another image](/new/path/images/2021/img.jpg)`,
		},
		{
			name:     "With Absolute Paths",
			markdown: `This is a test image: ![another image](/images/2021/img.jpg)`,
			expected: `This is a test image: ![another image](/images/2021/img.jpg)`,
		},
		{
			name:     "With Absolute URLs",
			markdown: `This is a test image: ![test image](http://example.com/images/test.png)`,
			expected: `This is a test image: ![test image](http://example.com/images/test.png)`,
		},
		{
			name:     "No Image Tags",
			markdown: `This is text with no images.`,
			expected: `This is text with no images.`,
		},
		{
			name:     "Invalid image syntax, no URL",
			markdown: `![invalid image]()`,
			expected: `![invalid image]()`,
		},
		{
			name:     "Invalid image syntax, no bucket",
			markdown: `![invalid image]`,
			expected: `![invalid image]`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := UpdateImagePaths([]byte(tt.markdown), "/new/path")
			assert.Equal(t, tt.expected, string(result), "output should match expected")
		})
	}
}

func TestUpdateFeaturedImagePaths(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		expected string
	}{
		{
			name:     "Relative Path",
			path:     "images/test.png",
			expected: "/new/path/images/test.png",
		},
		{
			name:     "With Absolute URL",
			path:     "http://example.com/images/test.png",
			expected: "http://example.com/images/test.png",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := UpdateFeaturedImagePaths("/new/path", tt.path)
			assert.Equal(t, tt.expected, result, "output should match expected")
		})
	}
}
