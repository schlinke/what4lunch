package urlparser

import (
	"testing"
	"time"
)

func TestParseURL(t *testing.T) {
	tn := time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)
	url := "https://testdomain.test/{year}/{intmonth}/{fridaydate}"

	v := ParseURL(url, tn)
	if v != "https://testdomain.test/2020/1/3.1." {
		t.Error("Expected 1, got ", v)
	}
}
