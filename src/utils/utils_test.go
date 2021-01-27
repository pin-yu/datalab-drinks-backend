package utils

import "testing"

func TestLastFridayNoon(t *testing.T) {
	lastFridayNoon := LastFridayNoon()
	if lastFridayNoon.Weekday() != 5 {
		t.Error("bad weekday in LastFridayNoon")
	}

	if lastFridayNoon.Hour() != 12 {
		t.Error("bad hour in LastFridayNoon")
	}
}

func TestDaysFromFriday(t *testing.T) {
	var subtractDay int

	subtractDay = daysFromFriday(5, 11)
	if subtractDay != -7 {
		t.Error("bad subtractDay")
	}

	subtractDay = daysFromFriday(5, 12)
	if subtractDay != 0 {
		t.Error("bad subtractDay")
	}

	subtractDay = daysFromFriday(6, 12)
	if subtractDay != -1 {
		t.Error("bad subtractDay")
	}

	subtractDay = daysFromFriday(4, 12)
	if subtractDay != -6 {
		t.Error("bad subtractDay")
	}
}
