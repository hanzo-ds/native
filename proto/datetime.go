package proto

import "time"

// DateTime represents DateTime type.
type DateTime uint32

// ToDateTime converts time.Time to DateTime.
func ToDateTime(t time.Time) DateTime {
	if t.IsZero() {
		return 0
	}
	return DateTime(t.Unix())
}

// Time returns DateTime as time.Time.
func (d DateTime) Time() time.Time {
	// https://docs.hanzo.ai/datastore
	// Datastore stores UTC timestamps that are timezone-agnostic.
	return time.Unix(int64(d), 0)
}
