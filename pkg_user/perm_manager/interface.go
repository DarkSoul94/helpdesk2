package perm_manager

type PermManager interface {
	ParseFromActions(actions ...string) error
	CheckPermissionForActions(permissions []byte, actions ...string) error
	CheckPermissionsGroup(permissions []byte, actions ...string) ([]string, error)
	ExportPermissionsList() ([]byte, error)
}
