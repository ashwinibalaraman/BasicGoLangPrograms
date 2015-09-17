package sleepTime

import (
	"testing"
	"time"
)

func TestSleepTime(t *testing.T) {

	startingTime := time.Now()
	endingTime := Sleep(time.Second * 6)

	t1 := startingTime.Second()

	t2 := endingTime.Second()

	var duration = t2 - t1

	if (duration) != 6 {
		t.Error("Expected", 6, " seconds sleep got ", duration, " seconds sleep")
	}
}
