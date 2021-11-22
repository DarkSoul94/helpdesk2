package internal_models

type TicketStatus struct {
	ID   uint64
	Name string
}

const (
	KeyTSWait           string = "ожидание"
	KeyTSResolve        string = "согласование"
	KeyTSInWork         string = "работа"
	KeyTSImplementation string = "реализация"
	KeyTSRevision       string = "доработка"
	KeyTSPostponed      string = "отложен"
	KeyTSRejected       string = "отклонен"
	KeyTSCompleted      string = "выполнен"
)

var TicketStatusMap map[string]TicketStatus = map[string]TicketStatus{
	KeyTSWait: {
		ID:   2,
		Name: "В ожидании",
	},
	KeyTSResolve: {
		ID:   3,
		Name: "В ожидании согласования",
	},
	KeyTSInWork: {
		ID:   4,
		Name: "В работе",
	},
	KeyTSImplementation: {
		ID:   5,
		Name: "В процессе реализации",
	},
	KeyTSRevision: {
		ID:   6,
		Name: "Отправлен на доработку",
	},
	KeyTSPostponed: {
		ID:   7,
		Name: "Отложен",
	},
	KeyTSRejected: {
		ID:   8,
		Name: "Отклонен",
	},
	KeyTSCompleted: {
		ID:   9,
		Name: "Выполнен",
	},
}

func (ts *TicketStatus) Set(Key string) {
	*ts = TicketStatusMap[Key]
}
