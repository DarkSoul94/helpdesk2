package perm_manager

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Manager struct {
	permissions PermLayer
}

func (m *Manager) CreateManagerFromActions(actions ...string) {
	m.permissions = PermLayer{
		FinalPerm:     make([]string, 0),
		SubPermGroups: make(map[string]PermLayer),
	}
	for _, action := range actions {
		permParts := strings.Split(action, ".")
		if len(permParts) == 1 {
			m.permissions.FinalPerm = append(m.permissions.FinalPerm, permParts[0])
		} else {
			m.putToLayer(&m.permissions, permParts)
		}
	}
}

func (m *Manager) formPermLayer(actions []string) PermLayer {
	var layer PermLayer = PermLayer{
		FinalPerm:     make([]string, 0),
		SubPermGroups: make(map[string]PermLayer),
	}

	if len(actions) != 1 {
		layer.SubPermGroups[actions[0]] = m.formPermLayer(actions[1:])
	} else {
		layer.FinalPerm = append(layer.FinalPerm, actions[0])
	}

	return layer
}

func (m *Manager) putToLayer(targetLayer *PermLayer, actions []string) {
	if len(actions) != 1 {
		if _, ok := targetLayer.SubPermGroups[actions[0]]; !ok {
			targetLayer.SubPermGroups[actions[0]] = m.formPermLayer(actions[1:])
		} else {
			pointer := targetLayer.SubPermGroups[actions[0]]
			m.putToLayer(&pointer, actions[1:])
			targetLayer.SubPermGroups[actions[0]] = pointer
		}
	} else {
		targetLayer.FinalPerm = append(targetLayer.FinalPerm, actions[0])
	}
}

func (m *Manager) ExportPermissionsList() (PermLayer, error) {
	return m.permissions, nil
}

func (m *Manager) CheckPermissionForActions(permissions []byte, actions ...string) error {
	var user_perm PermLayer
	err := json.Unmarshal(permissions, &user_perm)
	if err != nil {
		return err
	}
	if len(actions) != 0 {
		for _, action := range actions {
			keys := strings.Split(action, ".")
			err = m.checkInSlice(user_perm, keys)
			if err != nil {
				return fmt.Errorf(ErrNotAlowed.Error(), action)
			}
		}
	} else {
		return ErrNoneTargetActions
	}
	return nil
}

func (m *Manager) checkInSlice(layer PermLayer, keys []string) error {
	var err error
	if len(keys) != 1 {
		if _, ok := layer.SubPermGroups[keys[0]]; !ok {
			return ErrNotAlowed
		}

		err = m.checkInSlice(layer.SubPermGroups[keys[0]], keys[1:])
		if err != nil {
			return err
		}
	} else {
		for _, val := range layer.FinalPerm {
			if val == keys[0] {
				return nil
			}
		}
		return ErrNotAlowed
	}
	return nil
}

func (m *Manager) CheckPermissionsGroup(permissions []byte, actions ...string) ([]string, error) {
	var user_perm PermLayer
	out_actions := make([]string, 0)
	err := json.Unmarshal(permissions, &user_perm)
	if err != nil {
		return out_actions, err
	}
	for _, action := range actions {
		keys := strings.Split(action, ".")
		out_actions, err = m.checkInMap(m.permissions, keys)
		if err != nil {
			return out_actions, fmt.Errorf(ErrNotAlowed.Error(), err)
		}
	}
	return out_actions, nil
}

func (m *Manager) checkInMap(layer PermLayer, keys []string) ([]string, error) {
	var err error
	out := make([]string, 0)
	if _, ok := layer.SubPermGroups[keys[0]]; !ok {
		return out, ErrNotAlowed
	}
	if len(keys) != 1 {
		out, err = m.checkInMap(layer.SubPermGroups[keys[0]], keys[1:])
		if err != nil {
			return out, err
		}
	} else {
		temp := layer.SubPermGroups[keys[0]]
		out = append(out, temp.FinalPerm...)
	}
	return out, nil
}
