package utils

import (
	"testing"
	"time"
)

func TestOrderIntervalStartTime(t *testing.T) {
	const weekday = 5
	const drinksDeadlineHour = 16

	orderIntervalStartTime := OrderIntervalStartTime()
	if orderIntervalStartTime.Weekday() != weekday {
		t.Error("bad weekday in OrderIntervalStartTime")
	}

	if orderIntervalStartTime.Hour() != drinksDeadlineHour {
		t.Error("bad hour in OrderIntervalStartTime")
	}
}

func TestOrderIntervalEndTime(t *testing.T) {

	orderIntervalStartTime := OrderIntervalStartTime()
	orderIntervalEndTime := OrderIntervalEndTime()

	weekDuration, err := time.ParseDuration("168h")
	if err != nil {
		t.Error("bad ParseDuration in TestOrderIntervalEndTime")
	}

	if orderIntervalEndTime.Sub(orderIntervalStartTime) != weekDuration {
		t.Error("bad time in OrderIntervalEndTime")
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
