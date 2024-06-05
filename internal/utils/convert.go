package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

var countryToCode = map[string]string{
	"Thailand": "TH",
	"Japan":    "JP",
}

var englishShortMonths = map[time.Month]string{
	time.January:   "Jan",
	time.February:  "Feb",
	time.March:     "Mar",
	time.April:     "Apr",
	time.May:       "May",
	time.June:      "Jun",
	time.July:      "Jul",
	time.August:    "Aug",
	time.September: "Sep",
	time.October:   "Oct",
	time.November:  "Nov",
	time.December:  "Dec",
}

var thaiMonths = map[time.Month]string{
	time.January:   "มกราคม",
	time.February:  "กุมภาพันธ์",
	time.March:     "มีนาคม",
	time.April:     "เมษายน",
	time.May:       "พฤษภาคม",
	time.June:      "มิถุนายน",
	time.July:      "กรกฎาคม",
	time.August:    "สิงหาคม",
	time.September: "กันยายน",
	time.October:   "ตุลาคม",
	time.November:  "พฤศจิกายน",
	time.December:  "ธันวาคม",
}

var thaiWeekdays = map[time.Weekday]string{
	time.Sunday:    "วันอาทิตย์",
	time.Monday:    "วันจันทร์",
	time.Tuesday:   "วันอังคาร",
	time.Wednesday: "วันพุธ",
	time.Thursday:  "วันพฤหัสบดี",
	time.Friday:    "วันศุกร์",
	time.Saturday:  "วันเสาร์",
}

func ToDayMonth(t time.Time) string {
	return fmt.Sprintf("%02d %s", t.Day(), englishShortMonths[t.Month()])
}

func ToYearMonthDayRange(from, to time.Time) string {
	if from.Format("2016-01-02") == to.Format("2016-01-02") {
		return fmt.Sprintf("%sที่ %d %s %d", thaiWeekdays[from.Weekday()], from.Day(), thaiMonths[from.Month()], from.Year())
	}

	if from.Year() == to.Year() && from.Month() == to.Month() {
		return fmt.Sprintf("%sที่ %d - %sที่ %d %s %d", thaiWeekdays[from.Weekday()], from.Day(), thaiWeekdays[to.Weekday()], to.Day(), thaiMonths[from.Month()], from.Year())
	}

	if from.Year() == to.Year() {
		return fmt.Sprintf("%sที่ %d %s - %sที่ %d %s %d", thaiWeekdays[from.Weekday()], from.Day(), thaiMonths[from.Month()], thaiWeekdays[to.Weekday()], to.Day(), thaiMonths[to.Month()], from.Year())
	}

	return fmt.Sprintf("%sที่ %d %s %d - %sที่ %d %s %d", thaiWeekdays[from.Weekday()], from.Day(), thaiMonths[from.Month()], from.Year(), thaiWeekdays[to.Weekday()], to.Day(), thaiMonths[to.Month()], to.Year())
}

func ToFlagEmoji(country string) string {
	code, exists := countryToCode[country]
	if !exists {
		return country
	}

	var sb strings.Builder
	for _, rune := range code {
		emojiRune := 127397 + rune // 127397 is the offset for regional indicator symbols
		sb.WriteRune(emojiRune)
	}

	return sb.String()
}

func ToStruct(v map[string]interface{}, out interface{}) error {
	if v == nil {
		return errors.New("input map is nil")
	}

	bm, err := json.Marshal(v)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(bm, &out); err != nil {
		return err
	}

	return nil
}
