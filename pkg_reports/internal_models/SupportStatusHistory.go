package internal_models

import "time"

type SupportStatusHistory struct {
	StatusName string
	SelectTime time.Time
	Duration   time.Duration
}
