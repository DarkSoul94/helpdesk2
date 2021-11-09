package http

/*
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
*/
