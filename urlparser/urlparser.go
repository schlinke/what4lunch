package urlparser

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	dc "github.com/schlinke/what4lunch/datecalculator"
)

// ParseURL replaces the placesholder in the url with its correct values
// i.e. {year} to 2020
func ParseURL(url string, t time.Time) string {

	newurl := strings.Replace(url, "{year}", strconv.Itoa(dc.GetYear(t)), -1)
	newurl = strings.Replace(newurl, "{intmonth}", strconv.Itoa(dc.GetIntMonth(t)), -1)
	if strings.Contains(newurl, "{mondaydate}") {
		a, b := dc.GetDayAndMonth(dc.GetDayOfCW(t, 1))
		newurl = strings.Replace(newurl, "{mondaydate}", fmt.Sprintf("%d.%d.", a, b), -1)
	}
	if strings.Contains(newurl, "{fridaydate}") {
		a, b := dc.GetDayAndMonth(dc.GetDayOfCW(t, 5))
		newurl = strings.Replace(newurl, "{fridaydate}", fmt.Sprintf("%d.%d.", a, b), -1)
	}

	return newurl
}
