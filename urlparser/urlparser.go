package urlparser

import (
	"strconv"
	"strings"
	"time"
)

func getCW(t time.Time) int {
	_, week := t.ISOWeek()

	return week
}

func getYear(t time.Time) int {
	year := t.Year()

	return year
}

func getMonth(t time.Time) time.Month {
	month := t.Month()

	return month
}

func getIntMonth(t time.Time) int {
	var i int = int(getMonth(t))

	return i
}

// ParseURL replaces the placesholder in the url with its correct values
// i.e. {year} to 2020
func ParseURL(url string, t time.Time) string {

	newurl := strings.Replace(url, "{year}", strconv.Itoa(getYear(t)), -1)
	newurl = strings.Replace(newurl, "{intmonth}", strconv.Itoa(getIntMonth(t)), -1)

	return newurl
}
