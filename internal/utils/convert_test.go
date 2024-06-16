package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestConvertShowDayMonth(t *testing.T) {
	testTime := time.Date(2022, time.March, 15, 0, 0, 0, 0, time.UTC)
	expected := "15 Mar"
	result := ToDayMonth(testTime)
	assert.Equal(t, expected, result, "Date should be formatted as 'DD MMM'")
}

func TestConvertShowDate(t *testing.T) {
	tests := []struct {
		name     string
		from     time.Time
		to       time.Time
		expected string
	}{
		{
			name:     "same day",
			from:     time.Date(2023, time.April, 10, 0, 0, 0, 0, time.UTC),
			to:       time.Date(2023, time.April, 10, 0, 0, 0, 0, time.UTC),
			expected: "วันจันทร์ที่ 10 เมษายน 2023",
		},
		{
			name:     "same month, different days",
			from:     time.Date(2023, time.April, 10, 0, 0, 0, 0, time.UTC),
			to:       time.Date(2023, time.April, 15, 0, 0, 0, 0, time.UTC),
			expected: "วันจันทร์ที่ 10 - วันเสาร์ที่ 15 เมษายน 2023",
		},
		{
			name:     "different months, same year",
			from:     time.Date(2023, time.April, 30, 0, 0, 0, 0, time.UTC),
			to:       time.Date(2023, time.May, 1, 0, 0, 0, 0, time.UTC),
			expected: "วันอาทิตย์ที่ 30 เมษายน - วันจันทร์ที่ 1 พฤษภาคม 2023",
		},
		{
			name:     "different years",
			from:     time.Date(2022, time.December, 31, 0, 0, 0, 0, time.UTC),
			to:       time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC),
			expected: "วันเสาร์ที่ 31 ธันวาคม 2022 - วันอาทิตย์ที่ 1 มกราคม 2023",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToYearMonthDayRange(tt.from, tt.to)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestCountryToFlagEmoji(t *testing.T) {
	assert.Equal(t, "\U0001F1F9\U0001F1ED", ToFlagEmoji("Thailand"), "Should return Thai flag emoji")
	assert.Equal(t, "Mars", ToFlagEmoji("Mars"), "Should return the country name if not recognized")
}

func TestToStruct(t *testing.T) {
	type SimpleStruct struct {
		Name  string `json:"name"`
		Age   int    `json:"age"`
		Email string `json:"email"`
	}

	tests := []struct {
		name       string
		input      map[string]interface{}
		expected   SimpleStruct
		shouldFail bool
	}{
		{
			name: "Successful conversion",
			input: map[string]interface{}{
				"name":  "John Doe",
				"age":   30,
				"email": "john.doe@example.com",
			},
			expected: SimpleStruct{
				Name:  "John Doe",
				Age:   30,
				Email: "john.doe@example.com",
			},
			shouldFail: false,
		},
		{
			name: "Type mismatch",
			input: map[string]interface{}{
				"name":  123, // Should be a string
				"age":   30,
				"email": "john.doe@example.com",
			},
			expected:   SimpleStruct{},
			shouldFail: true,
		},
		{
			name: "Missing fields",
			input: map[string]interface{}{
				"name": "John Doe",
			},
			expected: SimpleStruct{
				Name: "John Doe",
			},
			shouldFail: false,
		},
		{
			name:       "Nil input map",
			input:      nil,
			expected:   SimpleStruct{},
			shouldFail: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result SimpleStruct
			err := ToStruct(tt.input, &result)

			if tt.shouldFail {
				assert.Error(t, err, "Expected an error but did not get one")
			} else {
				assert.NoError(t, err, "Did not expect an error but got one")
				assert.Equal(t, tt.expected, result, "Expected and actual structs do not match")
			}
		})
	}
}
