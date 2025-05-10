package timeutil

import (
	"log"
	"time"
)

const jst = "Asia/Tokyo"

type TimeStamped struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (ts *TimeStamped) SetCreatedAt(now time.Time) {
	ts.CreatedAt = now
}

func (ts *TimeStamped) SetUpdatedAt(now time.Time) {
	ts.UpdatedAt = now
}

func (ts *TimeStamped) SetTimeAt(now time.Time) {
	ts.SetCreatedAt(now)
	ts.SetUpdatedAt(now)
}

func NowJST() time.Time {
	return time.Now().In(loadLocation(jst))
}

func loadLocation(tz string) *time.Location {
	loc, err := time.LoadLocation(tz)
	if err != nil {
		log.Fatalf("failed to load location: %v", err)
	}

	return loc
}
