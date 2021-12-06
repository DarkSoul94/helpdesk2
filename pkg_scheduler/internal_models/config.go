package internal_models

import "github.com/shopspring/decimal"

const (
	ConfKey_GraceTime = `grace_time`
	ConfKey_Penalty   = `late_penalty`
)

type Config struct {
	GraceTime uint64
	Penalty   decimal.Decimal
}
