package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestOrderIntervalStartTime(t *testing.T) {
	const weekday = 5
	const drinksDeadlineHour = 16

	orderIntervalStartTime := OrderIntervalStartTime()
	assert.Equal(t, int(orderIntervalStartTime.Weekday()), weekday, "bad weekday in OrderIntervalStartTime")
	assert.Equal(t, int(orderIntervalStartTime.Hour()), drinksDeadlineHour, "bad hour in OrderIntervalStartTime")
}

func TestMeetingStartTime(t *testing.T) {
	meetingStartTime := MeetingStartTime()
	orderIntervalEndTime := OrderIntervalEndTime()

	assert.Equal(t, orderIntervalEndTime.Day(), meetingStartTime.Day(), "wrong meeting day")
	assert.Equal(t, 13, meetingStartTime.Hour(), "meeting start time is not at 13:00")
}

func TestOrderIntervalEndTime(t *testing.T) {
	orderIntervalStartTime := OrderIntervalStartTime()
	orderIntervalEndTime := OrderIntervalEndTime()

	// 24 hours * 7 days = 168 hours
	weekDuration, err := time.ParseDuration("168h")
	assert.Nil(t, err, "bad ParseDuration in TestOrderIntervalEndTime")
	assert.Equal(t, weekDuration, orderIntervalEndTime.Sub(orderIntervalStartTime), "bad time in OrderIntervalEndTime")
}

func TestHowManyDaysLastFromFridayToNow(t *testing.T) {
	var subtractDay int

	subtractDay = howManyDaysLastFromFridayToNow(5, 12)
	assert.Equal(t, -7, subtractDay, "bad subtractDay")

	subtractDay = howManyDaysLastFromFridayToNow(5, 16)
	assert.Equal(t, 0, subtractDay, "bad subtractDay")

	subtractDay = howManyDaysLastFromFridayToNow(6, 12)
	assert.Equal(t, -1, subtractDay, "bad subtractDay")

	subtractDay = howManyDaysLastFromFridayToNow(4, 16)
	assert.Equal(t, -6, subtractDay, "bad subtractDay")
}
