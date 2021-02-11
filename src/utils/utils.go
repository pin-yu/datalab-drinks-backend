package utils

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"runtime"
	"time"
)

// currently, meeting day is on Friday
const (
	meetingDay         = 5  // 5 means Friday
	meetingHour        = 13 // 13 means 13:00
	drinksDeadlineHour = 16 // 16 means 16:00, the order could be made until 16:00 Friday
	daysInWeek         = 7  // how many days in a week...
)

// ReadFile returns the content of file in a string
func ReadFile(file string) string {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}

// GetBasePath will return the project base path no matter where you run the program
func GetBasePath() string {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	return basepath
}

func lastFridayDate() (int, time.Month, int) {
	now := time.Now()

	hour := int(now.Hour())
	weekday := int(now.Weekday())

	subtractDay := howManyDaysLastFromFridayToNow(weekday, hour)
	// subtactDay is a negative time.
	lastFriday := now.AddDate(0, 0, subtractDay)

	// yy = years, mm = months, dd = days
	yy, mm, dd := lastFriday.Date()

	return yy, mm, dd
}

// OrderIntervalStartTime returns the start time of the order interval in time object
func OrderIntervalStartTime() time.Time {
	yy, mm, dd := lastFridayDate()

	orderIntervalStartTime := time.Date(yy, mm, dd, drinksDeadlineHour, 0, 0, 0, time.Local)

	return orderIntervalStartTime
}

// MeetingStartTime returns meeting start time
func MeetingStartTime() time.Time {
	yy, mm, dd := lastFridayDate()

	// the time will be the meeting time of this week after adding a week days
	meetingStartTime := time.Date(yy, mm, dd, meetingHour, 0, 0, 0, time.Local).AddDate(0, 0, daysInWeek)

	return meetingStartTime
}

// OrderIntervalEndTime returns the end time of the order interval in time object
func OrderIntervalEndTime() time.Time {
	startTime := OrderIntervalStartTime()

	const years = 0
	const months = 0
	const days = 7

	// simply add 7 days in to start time
	endTime := startTime.AddDate(years, months, days)

	return endTime
}

// calculate how many days last from last Friday to now
func howManyDaysLastFromFridayToNow(weekday int, hour int) int {
	var subtractDay int

	if weekday < meetingDay {
		subtractDay = weekday + (daysInWeek - meetingDay)
	} else {
		subtractDay = weekday % meetingDay
	}

	// if the current time is on Friday morning before 12 p.m.
	// compare it with the previous friday
	// else just compare it with that day
	if weekday == meetingDay && hour < drinksDeadlineHour {
		subtractDay = daysInWeek
	}

	return -subtractDay
}
