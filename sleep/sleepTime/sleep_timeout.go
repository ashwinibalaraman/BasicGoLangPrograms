package sleepTime

import "time"

func Sleep(n time.Duration) time.Time {
	return <-time.After(n)

}
