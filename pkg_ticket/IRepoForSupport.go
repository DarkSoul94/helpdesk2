package pkg_ticket

type IRepoForSupport interface {
	GetTicketsCount(supportID, statusID uint64) int
	GetTodayTicketsCount(supportID, statusID uint64) int
	Close() error
}
