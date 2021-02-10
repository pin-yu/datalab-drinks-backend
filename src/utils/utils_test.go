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

func TestOrderIntervalEndTime(t *testing.T) {
	orderIntervalStartTime := OrderIntervalStartTime()
	orderIntervalEndTime := OrderIntervalEndTime()

	// 24 hours * 7 days = 168 hours
	weekDuration, err := time.ParseDuration("168h")
	assert.Nil(t, err, "bad ParseDuration in TestOrderIntervalEndTime")
	assert.Equal(t, orderIntervalEndTime.Sub(orderIntervalStartTime), weekDuration, "bad time in OrderIntervalEndTime")
}

func TestDaysFromFriday(t *testing.T) {
	var subtractDay int

	subtractDay = daysFromFriday(5, 12)
	assert.Equal(t, subtractDay, -7, "bad subtractDay")

	subtractDay = daysFromFriday(5, 16)
	assert.Equal(t, subtractDay, 0, "bad subtractDay")

	subtractDay = daysFromFriday(6, 12)
	assert.Equal(t, subtractDay, -1, "bad subtractDay")

	subtractDay = daysFromFriday(4, 16)
	assert.Equal(t, subtractDay, -6, "bad subtractDay")
}
