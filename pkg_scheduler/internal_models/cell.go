package internal_models

import "time"

type Cell struct {
	ID        uint64
	SupportID uint64
	OfficeID  uint64
	StartTime time.Time
	EndTime   time.Time
	Date      time.Time
	Vacation  bool
	SickLeave bool
}

func (sc *Cell) Late() bool {
	var (
		now, shift checkTime
	)
	now.hour, now.min, _ = time.Now().Local().Clock()
	shift.hour, shift.min, _ = sc.StartTime.Clock()
	return now.compare(&shift)
}

func (sc *Cell) WasOpenedToday() bool {
	var (
		now, open checkTime
	)
	now.hour, now.min, _ = time.Now().Local().Clock()
	open.hour, open.min, _ = sc.StartTime.Clock()
	return now.compare(&open)
}

type checkTime struct {
	hour int
	min  int
}

func (t *checkTime) compare(shift *checkTime) bool {
	if t.hour > shift.hour {
		return true
	}
	if t.hour == shift.hour {
		if t.min > shift.min {
			return true
		}
	}
	return false
}
