package durations

import (
	"fmt"
	"testing"
)

func Test_GetDate(t *testing.T) {

	time, err := getdate("3mon 2w 1h")

	if err != nil {
		t.Error(err.Error())
	}
	if 15 > time {
		t.Error("error")
	}
	fmt.Println(time)
}

func Test_parsetime(t *testing.T) {

	time, err := parsetime("2", "w")
	if err != nil {
		t.Error(err.Error())
	}
	if 15 > time {
		t.Error("error")
	}
}
