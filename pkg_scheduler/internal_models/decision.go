package internal_models

import (
	"sort"

	"github.com/DarkSoul94/helpdesk2/models"
)

type Decision struct {
	ID    uint64
	Text  string
	valid bool
	value bool
}

var (
	decisionNone Decision = Decision{
		ID:    0,
		Text:  "",
		valid: false,
	}
	decisionPardon Decision = Decision{
		ID:    1,
		Text:  "Помилован",
		valid: true,
		value: true,
	}
	decisionExecute Decision = Decision{
		ID:    2,
		Text:  "Казнен",
		valid: true,
		value: false,
	}

	decisionMap map[uint64]Decision = map[uint64]Decision{
		1: decisionPardon,
		2: decisionExecute,
	}
)

func GetLateDecision(id uint64) (Decision, models.Err) {
	if val, ok := decisionMap[id]; ok {
		return val, nil
	}
	return Decision{}, models.BadRequest("Decision with this ID does not exist")
}

func SetLateDecision(valid, decision bool) Decision {
	if !valid {
		return decisionNone
	}
	if decision {
		return decisionPardon
	}
	return decisionExecute
}

func (ld *Decision) GetDecisionValue() (valid, value bool) {
	return decisionMap[ld.ID].valid, decisionMap[ld.ID].value
}

func GetDicisionsList() []Decision {
	keys := make([]int, 0)
	outList := make([]Decision, 0)
	for key := range decisionMap {
		keys = append(keys, int(key))
	}
	sort.Ints(keys)
	for _, key := range keys {
		outList = append(outList, decisionMap[uint64(key)])
	}
	return outList
}
