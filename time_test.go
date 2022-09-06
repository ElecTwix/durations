package durations

import (
	"strconv"
	"testing"
	"time"
)

func Test_GetDate(t *testing.T) {

	strArr := []string{"c", "y", "h"}
	mulArr := []int{1, 1, 1}
	sumStr := ""

	var expected time.Duration = 0
	for index, str := range strArr {

		digit := strconv.Itoa(mulArr[index])
		expected += time.Duration(unitMap[str] * uint64(mulArr[index]))
		sumStr += str + digit + " "
		if str == "c" {
			expected += GetTotalLeapYear(mulArr[index] * 100)
		}
		if str == "y" {
			expected += GetTotalLeapYear(mulArr[index])
		}

	}

	time, err := GetDuration(sumStr)

	if err != nil {
		t.Error(err.Error())
	}
	if expected != time {
		t.Error("not same with expected value.")
	}
}

func Test_parsetime(t *testing.T) {

	str := "w"
	mul := 3

	var expected time.Duration = 0

	expected = expected + time.Duration(unitMap[str]*uint64(mul))
	mulstr := strconv.Itoa(mul)
	time, err := parseTime(mulstr, str)
	if err != nil {
		t.Error(err.Error())
	}
	if expected != time {
		t.Error("not same with expected value.")
	}
}
