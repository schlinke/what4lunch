package datecalculator

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

	v := GetYear(tn)
	if v != 2020 {
		t.Error("Expected 2020, got ", v)
	}
}

func TestGetDay(t *testing.T) {
	tn := time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)

	v := getDay(tn)
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

	v := GetIntMonth(tn)
	if v != 1 {
		t.Error("Expected 1, got ", v)
	}
}

func TestGetDayAndMonth(t *testing.T) {
	tn := time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)

	v1, v2 := GetDayAndMonth(tn)
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
	v1 := GetDayOfCW(tn, 1)
	if v1 != time.Date(2019, time.December, 30, 0, 0, 0, 0, time.UTC) {
		t.Error("Expected 2019-12-30 00:00:00 +0000 UTC, got ", v1)
	}

	// Test for friday
	v2 := GetDayOfCW(tn, 5)
	if v2 != time.Date(2020, time.January, 3, 0, 0, 0, 0, time.UTC) {
		t.Error("Expected 2020-01-03 00:00:00 +0000 UTC, got ", v2)
	}
}
