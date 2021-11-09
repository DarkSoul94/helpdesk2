package group_manager

type Repo interface {
	GetGroupByID(groupID uint64) (*Group, error)
	GetGroupList() ([]*Group, error)

	Close() error
}
