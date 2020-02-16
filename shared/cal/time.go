package cal

import (
	"time"
)

func DateNow() string {
	t := time.Now()
	return t.Format(time.RFC1123)[:16]
}

func TimeNow() string {
	t := time.Now()
	return t.Format(time.Kitchen)
}

func ParseDuration(s string) (d time.Duration, err error) {
	d, err = time.ParseDuration(s)
	return
}

func Wait(d time.Duration) {
	time.Sleep(d)
}
