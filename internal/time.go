package internal

import "time"

var Tz, _ = time.LoadLocation("Asia/Shanghai")

func GetNowTime() time.Time {
	return time.Now().In(Tz)
}

func GetCalculateTime(currentTime time.Time, d string) (time.Time, error) {
	duration, err := time.ParseDuration(d)
	if err != nil {
		return time.Time{}, err
	}
	return currentTime.In(Tz).Add(duration), nil
}
