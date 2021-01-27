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
	drinksDeadlineHour = 12
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

// LastFridayNoon returns the time object of the last friday noon
func LastFridayNoon() *time.Time {
	now := time.Now()

	hour := int(now.Hour())
	weekday := int(now.Weekday())

	subtractDay := daysFromFriday(weekday, hour)
	lastFriday := now.AddDate(0, 0, subtractDay)

	yy, mm, dd := lastFriday.Date()
	lastFridayNoon := time.Date(yy, mm, dd, drinksDeadlineHour, 0, 0, 0, time.Local)

	return &lastFridayNoon
}

// calculate the day numbers from last friday to now
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
