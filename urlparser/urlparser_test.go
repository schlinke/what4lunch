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

func TestParseURL(t *testing.T) {
	tn := time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)
	url := "https://testdomain.test/{year}/{intmonth}"

	v := ParseURL(url, tn)
	if v != "https://testdomain.test/2020/1" {
		t.Error("Expected 1, got ", v)
	}
}
