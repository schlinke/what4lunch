package urlparser

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func getCW(t time.Time) int {
	_, week := t.ISOWeek()

	return week
}

func getDay(t time.Time) int {
	day := t.Day()

	return day
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

// Return the date of a given Weekday i.e. Monday of the current calenderweek
// day parameter as describe in Weekday doc (Sunday = 0, ....)
func getDayOfCW(t time.Time, day int) time.Time {
	daycurrent := t.Weekday()
	dayint := int(daycurrent)
	difference := time.Duration(day - dayint)

	newday := t.Add(time.Hour * 24 * difference)

	return newday
}

func getDayAndMonth(t time.Time) (day, month int) {
	day = getDay(t)
	month = getIntMonth(t)

	return day, month
}

// ParseURL replaces the placesholder in the url with its correct values
// i.e. {year} to 2020
func ParseURL(url string, t time.Time) string {

	newurl := strings.Replace(url, "{year}", strconv.Itoa(getYear(t)), -1)
	newurl = strings.Replace(newurl, "{intmonth}", strconv.Itoa(getIntMonth(t)), -1)
	if strings.Contains(newurl, "{mondaydate}") {
		a, b := getDayAndMonth(getDayOfCW(t, 1))
		newurl = strings.Replace(newurl, "{mondaydate}", fmt.Sprintf("%d.%d.", a, b), -1)
	}
	if strings.Contains(newurl, "{fridaydate}") {
		a, b := getDayAndMonth(getDayOfCW(t, 5))
		newurl = strings.Replace(newurl, "{fridaydate}", fmt.Sprintf("%d.%d.", a, b), -1)
	}

	return newurl
}
