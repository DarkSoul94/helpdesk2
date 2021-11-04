package http

import (
	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/DarkSoul94/helpdesk2/user_manager/perm_manager"
)

func (h *Handler) toOutUser(user models.User) outUser {
	group := h.toOutGroup(*user.Group)

	return outUser{
		ID:    user.ID,
		Name:  user.Name,
		Group: &group,
	}
}

func (h *Handler) toOutUserForList(user models.User) outUserForList {
	return outUserForList{
		ID:         user.ID,
		Name:       user.Name,
		GroupID:    user.Group.ID,
		Department: user.Department,
	}
}

func (h *Handler) toOutGroup(group models.Group) outGroup {
	return outGroup{
		ID:          group.ID,
		Name:        group.Name,
		Permissions: h.toOutPermissions(group.Permissions),
	}
}

func (h *Handler) toOutGroupForList(group models.Group) outGroup {
	return outGroup{
		ID:   group.ID,
		Name: group.Name,
	}
}

func (h *Handler) toOutPermissions(perm perm_manager.PermLayer) map[string]interface{} {
	out := make(map[string]interface{})
	if len(perm.SubPermGroups) != 0 {
		for key := range perm.SubPermGroups {
			out[key] = h.toOutPermissions(perm.SubPermGroups[key])
		}
	}
	if len(perm.FinalPerm) != 0 {
		temp := make([]string, 0)
		temp = append(temp, perm.FinalPerm...)
		out["actions"] = temp
	}
	return out
}
