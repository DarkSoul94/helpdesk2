package internal_models

type TicketStatus struct {
	ID   uint64
	Name string
}

const (
	KeyTSWait           string = "wait"
	KeyTSResolve        string = "resolve"
	KeyTSInWork         string = "in_work"
	KeyTSImplementation string = "implementation"
	KeyTSRevision       string = "revision"
	KeyTSPostponed      string = "postponed"
	KeyTSRejected       string = "rejected"
	KeyTSCompleted      string = "completed"

	TSWaitID           uint64 = 2
	TSWaitForResolveID uint64 = 3
	TSInWorkID         uint64 = 4
	TSImplementationID uint64 = 5
	TSRevisionID       uint64 = 6
	TSPostponedID      uint64 = 7
	TSRejectedID       uint64 = 8
	TSCompletedID      uint64 = 9
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
