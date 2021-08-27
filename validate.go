package unaswrappergo

import "time"

func tokenExpired(t time.Time) bool {
	now := time.Now()
	return t.After(now)

}
