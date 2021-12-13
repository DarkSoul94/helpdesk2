package internal_models

import "time"

type Shift struct {
	ID            uint64
	Support       *Support
	OpeningTime   time.Time
	ClosingTime   time.Time
	ClosingStatus bool
}

func (s *Shift) Reopen() {
	s.ClosingTime = time.Time{}
	s.ClosingStatus = false
}

func (s *Shift) Close() {
	var status Status
	status.notAccept()
	s.ClosingTime = time.Now()
	s.ClosingStatus = true
	s.Support.Status = &status
}

func (s *Shift) Open(supportID uint64, oTime time.Time) {
	var status Status
	status.accept()
	*s = Shift{
		Support: &Support{
			ID:     supportID,
			Status: &status,
		},
		OpeningTime: oTime,
	}
}

func (sc *Shift) WasOpenedToday() bool {
	var (
		now, open checkDate
	)
	now.year, now.month, now.day = time.Now().Local().Date()
	open.year, open.month, open.day = sc.OpeningTime.Date()
	return now.compare(&open)
}

type checkDate struct {
	year  int
	month time.Month
	day   int
}

func (c *checkDate) compare(val *checkDate) bool {
	if c.year == val.year &&
		c.month == val.month &&
		c.day == val.day {
		return true
	}
	return false
}
