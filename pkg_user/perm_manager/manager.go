package perm_manager

import (
	"encoding/json"
	"reflect"
	"strings"
)

type Manager struct {
	permissions map[string]layer
}

func NewManager(actions interface{}) (*Manager, error) {
	manager := &Manager{
		permissions: make(map[string]layer),
	}
	switch reflect.TypeOf(actions) {
	case reflect.TypeOf([]string{}):
		for _, action := range actions.([]string) {
			actionParts := strings.Split(action, ".")
			manager.init(actionParts, actionParts)
		}
	case reflect.TypeOf(map[string]string{}):
		for name, action := range actions.(map[string]string) {
			actionParts := strings.Split(action, ".")
			nameParts := strings.Split(name, ".")
			manager.init(actionParts, nameParts)
		}
	default:
		return nil, ErrWrongType
	}
	return manager, nil
}

func (m *Manager) init(actionParts, nameParts []string) {
	var temp layer

	if len(actionParts) != 1 {
		if _, ok := m.permissions[actionParts[0]]; !ok {
			temp.formLayer(actionParts[1:], nameParts)
			m.permissions[actionParts[0]] = temp
		} else {
			temp = m.permissions[actionParts[0]]
			temp.putToLayer(actionParts[1:], nameParts)
			m.permissions[actionParts[0]] = temp
		}
	} else {
		temp = m.permissions[KeyFinalActoins]
		temp.Actions = append(temp.Actions, targetAction{Name: actionParts[0], Action: nameParts[0]})
		m.permissions[KeyFinalActoins] = temp
	}
}

func (m *Manager) ExportToTreeView() []byte {
	tree := make([]treeLayer, 0)
	for key, val := range m.permissions {
		tree = append(tree, val.toTreeView(val.Name, key))
	}
	out, _ := json.Marshal(tree)
	return out
}
