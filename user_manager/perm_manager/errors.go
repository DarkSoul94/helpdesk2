package perm_manager

import "errors"

var (
	ErrNotAlowed         = errors.New("The action %s are not alowed")
	ErrNoneTargetActions = errors.New("List of targets actions is empty")
)
