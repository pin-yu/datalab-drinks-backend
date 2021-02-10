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
	meetingDay         = 5
	drinksDeadlineHour = 16
	daysInWeek         = 7
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

// OrderIntervalStartTime returns the start time of the order interval in time object
func OrderIntervalStartTime() time.Time {
	now := time.Now()

	hour := int(now.Hour())
	weekday := int(now.Weekday())

	subtractDay := daysFromFriday(weekday, hour)
	lastFriday := now.AddDate(0, 0, subtractDay)

	yy, mm, dd := lastFriday.Date()
	lastFridayNoon := time.Date(yy, mm, dd, drinksDeadlineHour, 0, 0, 0, time.Local)

	return lastFridayNoon
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

// calculate how many number of days from last friday to now
func daysFromFriday(weekday int, hour int) int {
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
