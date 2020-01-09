package utils

import "time"

func GetYYYYMMDDHHMMSS() string {
	now := time.Now()
	return now.Format("20060102_150405")
}

func GetYYYYMMDD() string {
	now := time.Now()
	return now.Format("20060102")
}

type Escape struct {
	Day    int64
	Hour   int64
	Minute int64
	Second int64
}

func GetEscapeTime(escape int64) Escape {

	second := escape % 60
	minute := escape / 60 % 60
	hour := escape / (60 * 60) % 24
	day := escape / (60 * 60 * 24)

	return Escape{
		day, hour, minute, second,
	}
}
