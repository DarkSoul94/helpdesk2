package internal_models

type Support struct {
	ID       uint64
	UserID   uint64
	Status   *Status
	Priority bool
}
