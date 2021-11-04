package perm_manager

type PermLayer struct {
	FinalPerm     []string             `json:"actions_list,omitempty"`
	SubPermGroups map[string]PermLayer `json:"sub_perm,omitempty"`
}
