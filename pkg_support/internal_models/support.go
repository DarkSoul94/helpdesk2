package internal_models

type Support struct {
	SupportID uint64
	Status    *Status
	Priority  bool
}

//NewSupport создает объект нового саппорта с установленным статусом "Не принимаю запросы"
func NewSupport(userID uint64) *Support {
	return &Support{
		SupportID: userID,
		Priority:  false,
		Status: &Status{
			ID: 4,
		},
	}
}
