package http

import (
	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/DarkSoul94/helpdesk2/pkg_user/group_manager"
)

func (h *Handler) toInpGroup(group *models.Group) *inpGroup {
	return &inpGroup{
		ID:          group.ID,
		Name:        group.Name,
		Permissions: h.toOutPermissions(group.Permissions),
	}
}

func (h *Handler) toInpUser(user models.User, token string) inpUser {
	return inpUser{
		ID:    user.ID,
		Name:  user.Name,
		Token: token,
		Group: h.toInpGroup(user.Group),
	}
}

func (h *Handler) toOutPermissions(perm group_manager.PermLayer) map[string]interface{} {
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
