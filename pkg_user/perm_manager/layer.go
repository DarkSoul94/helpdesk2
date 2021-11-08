package perm_manager

const (
	KeyFinalActoins = "final actions"
)

type treeLayer struct {
	ID       string      `json:"id,omitempty"`
	Name     string      `json:"name,omitempty"`
	Children []treeLayer `json:"children,omitempty"`
}

type targetAction struct {
	Name   string `json:"name,omitempty"`
	Action string `json:"action,omitempty"`
}

type layer struct {
	Name     string           `json:"name,omitempty"`
	SubLayer map[string]layer `json:"sub_layer,omitempty"`
	Actions  []targetAction   `json:"actions,omitempty"`
}

func (l *layer) formLayer(actions, name []string) {
	l.Name = name[0]
	l.SubLayer = make(map[string]layer)

	var sub layer = layer{
		Name:     "",
		SubLayer: make(map[string]layer),
		Actions:  make([]targetAction, 0),
	}

	if len(actions) != 1 {
		sub.formLayer(actions[1:], name[1:])
		l.SubLayer[actions[0]] = sub
	} else {
		l.Actions = append(l.Actions, targetAction{Name: name[1], Action: actions[0]})
	}
}

func (l *layer) putToLayer(actions, name []string) {
	var temp layer
	if len(actions) != 1 {
		if _, ok := l.SubLayer[actions[0]]; !ok {
			temp.formLayer(actions[1:], name[1:])
			l.SubLayer[actions[0]] = temp
		} else {
			temp = l.SubLayer[actions[0]]
			temp.putToLayer(actions[1:], name[1:])
			l.SubLayer[actions[0]] = temp
		}
	} else {
		l.Actions = append(l.Actions, targetAction{Name: name[1], Action: actions[0]})
	}
}

func (l *layer) toTreeView(name, id string) treeLayer {
	var tLayer treeLayer = treeLayer{
		ID:       id,
		Name:     name,
		Children: make([]treeLayer, 0),
	}
	if len(l.SubLayer) != 0 {
		for key, val := range l.SubLayer {
			tID := id + "." + key
			tLayer.Children = append(tLayer.Children, val.toTreeView(val.Name, tID))
		}
	}
	if len(l.Actions) != 0 {
		for _, val := range l.Actions {
			tID := id + "." + val.Action
			tLayer.Children = append(tLayer.Children, treeLayer{ID: tID, Name: val.Name})
		}
	}
	return tLayer
}
