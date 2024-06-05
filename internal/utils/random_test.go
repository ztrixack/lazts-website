package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomizeBlackholes(t *testing.T) {
	testCases := []struct {
		name      string
		count     int
		expectLen int
	}{
		{name: "Positive count", count: 10, expectLen: 10},
		{name: "Zero count", count: 0, expectLen: 0},
		{name: "Negative count", count: -5, expectLen: 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			blackholes := RandomizeBlackholes(tc.count)
			assert.Equal(t, tc.expectLen, len(blackholes), "The number of blackholes generated should match the expected count")

			for _, bh := range blackholes {
				assert.True(t, bh.Size >= 180 && bh.Size <= 540, "Size should be within the range 180 to 540")
				assert.True(t, bh.Rotate >= 0 && bh.Rotate < 360, "Rotate should be within the range 0 to 359")
				assert.True(t, bh.Opacity >= 5 && bh.Opacity <= 110, "Opacity should be within the range 5 to 110")
				assert.True(t, bh.Width >= 15 && bh.Width <= 75, "Width should be within the range 15 to 75")
			}
		})
	}
}

func TestRandomizeClouds(t *testing.T) {
	testCases := []struct {
		name      string
		count     int
		expectLen int
	}{
		{name: "Positive count", count: 10, expectLen: 10},
		{name: "Zero count", count: 0, expectLen: 0},
		{name: "Negative count", count: -5, expectLen: 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			clouds := RandomizeClouds(tc.count)
			assert.Equal(t, tc.expectLen, len(clouds), "The number of clouds generated should match the expected count")

			for _, c := range clouds {
				assert.True(t, c.Top >= -50 && c.Top < 100, "Top should be within the range -50 to 99")
				assert.True(t, c.Left >= 0 && c.Left < 100, "Left should be within the range 0 to 99")
				assert.True(t, c.Rotate >= 0 && c.Rotate < 360, "Rotate should be within the range 0 to 359")
			}
		})
	}
}
