package datecalculator

import (
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

// GetYear returns the Year of a given date
func GetYear(t time.Time) int {
	year := t.Year()

	return year
}

func getMonth(t time.Time) time.Month {
	month := t.Month()

	return month
}

// GetIntMonth return the month of a given date as an integer
func GetIntMonth(t time.Time) int {
	var i int = int(getMonth(t))

	return i
}

// GetDayOfCW returns the date of a given Weekday i.e. Monday of the current calenderweek
// day parameter as describe in Weekday doc (Sunday = 0, ....)
func GetDayOfCW(t time.Time, day int) time.Time {
	daycurrent := t.Weekday()
	dayint := int(daycurrent)
	difference := time.Duration(day - dayint)

	newday := t.Add(time.Hour * 24 * difference)

	return newday
}

// GetDayAndMonth returns the day and month of a given date
func GetDayAndMonth(t time.Time) (day, month int) {
	day = getDay(t)
	month = GetIntMonth(t)

	return day, month
}
