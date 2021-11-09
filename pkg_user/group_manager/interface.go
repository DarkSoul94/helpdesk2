package group_manager

type PermManager interface {
	ExportToTreeView() []byte

	GetGroupByID(groupID uint64) (*Group, error)
	GetGroupList() ([]*Group, error)
}
