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
	Month                = 4 * Week
	Year                 = 12 * Week
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

func getdate(durationstr string) (totalduration time.Duration, err error) {

	strarr := strings.Fields(durationstr)

	var sumduration time.Duration = 0

	for _, str := range strarr {
		var intstr string = ""
		var charstr string = ""
		for _, char := range str {
			if '0' <= char && char <= '9' {
				intstr = intstr + string(char)
			} else {
				if !unicode.IsDigit(char) {
					charstr = charstr + string(char)
				}
			}
		}
		tempdur, err := parsetime(intstr, charstr)
		sumduration = sumduration + tempdur
		if err != nil {
			return 0, err
		}

	}

	return sumduration, nil
}

func parsetime(intstr string, charstr string) (duration time.Duration, err error) {
	digit, err := strconv.Atoi(intstr)

	dur := unitMap[charstr] * uint64(digit)

	return time.Duration(dur), nil
}
