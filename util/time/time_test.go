package time_test

import (
	. "fehu/util/time"
	"testing"
)

func Test1(t *testing.T) {

	a := GetWeekStartUnix()
	if a != 1615132800 {
		t.Error("err")
	}

}
