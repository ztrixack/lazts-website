package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateImagePaths(t *testing.T) {
	testcases := []struct {
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

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			result := UpdateImagePaths([]byte(tc.markdown), "/new/path")
			assert.Equal(t, tc.expected, string(result), "output should match expected")
		})
	}
}

func TestUpdateFeaturedImagePaths(t *testing.T) {
	testcases := []struct {
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

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			result := UpdateFeaturedImagePaths("/new/path", tc.path)
			assert.Equal(t, tc.expected, result, "output should match expected")
		})
	}
}
