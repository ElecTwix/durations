package durations

import (
	"strconv"
	"testing"
	"time"
)

func Test_GetDate(t *testing.T) {

	strArr := []string{"mon", "w", "h"}
	mulArr := []int{3, 2, 1}
	sumStr := ""

	var expected time.Duration = 0
	for index, val := range strArr {
		expected = expected + time.Duration(unitMap[val]*uint64(mulArr[index]))
		sumStr = sumStr + val + strconv.Itoa(mulArr[index]) + " "
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
