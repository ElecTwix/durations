package durations

import (
	"testing"
)

func Test_GetDate(t *testing.T) {

	time, err := GetDuration("3mon 2w 1h")

	if err != nil {
		t.Error(err.Error())
	}
	if 15 > time {
		t.Error("error")
	}
}

func Test_parsetime(t *testing.T) {

	time, err := parseTime("2", "w")
	if err != nil {
		t.Error(err.Error())
	}
	if 15 > time {
		t.Error("error")
	}
}
