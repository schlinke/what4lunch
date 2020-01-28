package urlparser

import (
	"testing"
	"time"
)

func TestGetCW(t *testing.T) {
	tn := time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)

	v := getCW(tn)
	if v != 1 {
		t.Error("Expected 1, got ", v)
	}
}

func TestGetYear(t *testing.T) {
	tn := time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)

	v := getYear(tn)
	if v != 2020 {
		t.Error("Expected 2020, got ", v)
	}
}

func TestGetDay(t *testing.T) {
	tn := time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)

	v := getYear(tn)
	if v != 1 {
		t.Error("Expected 1, got ", v)
	}
}

func TestGetMonth(t *testing.T) {
	tn := time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)

	v := getMonth(tn)
	if v != time.January {
		t.Error("Expected January, got ", v)
	}
}

func TestGetIntMonth(t *testing.T) {
	tn := time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)

	v := getIntMonth(tn)
	if v != 1 {
		t.Error("Expected 1, got ", v)
	}
}

func TestGetDayAndMonth(t *testing.T) {
	tn := time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)

	v1, v2 := getDayAndMonth(tn)
	if v1 != 1 {
		t.Error("Expected 1, got ", v1)
	}
	if v1 != 1 {
		t.Error("Expected 1, got ", v2)
	}
}

func TestGetDayOfCW(t *testing.T) {
	// The date is Wednesday
	tn := time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)

	// Test for monday
	v1 := getDayOfCW(tn, 1)
	if v1 != time.Date(2019, time.December, 30, 0, 0, 0, 0, time.UTC) {
		t.Error("Expected 2019-12-30 00:00:00 +0000 UTC, got ", v1)
	}

	// Test for friday
	v2 := getDayOfCW(tn, 5)
	if v2 != time.Date(2020, time.January, 4, 0, 0, 0, 0, time.UTC) {
		t.Error("Expected 2020-01-03 00:00:00 +0000 UTC, got ", v2)
	}
}

func TestParseURL(t *testing.T) {
	tn := time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)
	url := "https://testdomain.test/{year}/{intmonth}"

	v := ParseURL(url, tn)
	if v != "https://testdomain.test/2020/1" {
		t.Error("Expected 1, got ", v)
	}
}
