package internal

import "time"

func timeMst() string {
	return time.Now().Format("2006-01-02 15:04:05 MST")
}
