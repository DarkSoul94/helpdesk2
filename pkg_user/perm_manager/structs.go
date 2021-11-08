package perm_manager

type Group struct {
	Id   uint64
	Name string
	tree treeLayer
}
