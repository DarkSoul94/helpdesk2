package internal_models

type Support struct {
	ID       uint64
	Name     string
	Status   *Status
	Priority bool
}

//NewSupport создает объект нового саппорта с установленным статусом "Не принимаю запросы"
func NewSupport(userID uint64) *Support {
	return &Support{
		ID:       userID,
		Priority: false,
		Status: &Status{
			ID: 4,
		},
	}
}
