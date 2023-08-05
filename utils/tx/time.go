package tx

import "time"

func Now() time.Time {
	return time.Now()
}

func NowP() *time.Time {
	t := time.Now()
	return &t
}
