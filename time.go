package durations

import (
	"strconv"
	"strings"
	"time"
	"unicode"
)

type Duration int64

const (
	Nanosecond  Duration = 1
	Microsecond          = 1000 * Nanosecond
	Millisecond          = 1000 * Microsecond
	Second               = 1000 * Millisecond
	Minute               = 60 * Second
	Hour                 = 60 * Minute
	Day                  = 24 * Hour
	Week                 = 7 * Day
	Month                = 30 * Day
	Year                 = 365 * Day
	Century              = 100 * Year
)

var unitMap = map[string]uint64{
	"ns":  uint64(Nanosecond),
	"us":  uint64(Microsecond),
	"µs":  uint64(Microsecond), // U+00B5 = micro symbol
	"μs":  uint64(Microsecond), // U+03BC = Greek letter mu
	"ms":  uint64(Millisecond),
	"s":   uint64(Second),
	"m":   uint64(Minute),
	"h":   uint64(Hour),
	"d":   uint64(Day),
	"w":   uint64(Week),
	"mon": uint64(Month),
	"y":   uint64(Year),
	"c":   uint64(Century),
}

func GetDuration(durationstr string) (totalduration time.Duration, err error) {

	strarr := strings.Fields(durationstr)

	var sumduration time.Duration = 0

	for _, str := range strarr {
		var intstr string = ""
		var charstr string = ""
		for _, char := range str {
			if '0' <= char && char <= '9' {
				intstr += string(char)
			} else {
				if !unicode.IsDigit(char) {
					charstr = charstr + string(char)
				}
			}
		}
		tempdur, err := parseTime(intstr, charstr)
		sumduration += tempdur
		if err != nil {
			return 0, err
		}

	}

	return sumduration, nil
}

func parseTime(intstr string, charstr string) (duration time.Duration, err error) {
	var totalduration time.Duration = 0
	digit, err := strconv.Atoi(intstr)

	switch charstr {
	case "c":
		totalduration += GetTotalLeapYear(digit * 100)
	case "y":
		totalduration += GetTotalLeapYear(digit)
	}

	totalduration += time.Duration(unitMap[charstr] * uint64(digit))

	return totalduration, nil
}

func GetTotalLeapYear(totalyear int) (duration time.Duration) {
	var sumduration time.Duration = 0
	yearNow := time.Now().Year()
	for i := 0; totalyear > i; i++ {
		if isLeapYear(yearNow) {
			sumduration += time.Duration(Day)
		}
		yearNow++
	}
	return sumduration
}

func isLeapYear(year int) bool {
	if year%4 == 0 && year%100 != 0 || year%400 == 0 {
		return true
	}
	return false
}
