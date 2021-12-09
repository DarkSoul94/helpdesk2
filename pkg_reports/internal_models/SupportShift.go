package internal_models

type SupportsShifts struct {
	Support          string
	WithOutGraceTime uint64
	ShiftsCount      int
	MinutesCount     uint64
	DayTime          []OpeningDayTime
}

type OpeningDayTime struct {
	OpeningDate        string
	ClosingDate        string
	CountOfMinutesLate uint64
}
