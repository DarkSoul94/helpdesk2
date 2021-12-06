package http

import (
	"net/http"

	"github.com/DarkSoul94/helpdesk2/global_const"
	"github.com/DarkSoul94/helpdesk2/global_const/actions"
	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/DarkSoul94/helpdesk2/pkg_user/group_manager"
	"github.com/gin-gonic/gin"
)

type PermissionMiddleware struct {
	usecase group_manager.IPermManager
}

func NewPermissionMiddleware(usecase group_manager.IPermManager) gin.HandlerFunc {
	return (&PermissionMiddleware{
		usecase: usecase,
	}).CheckPermissions
}

func (p *PermissionMiddleware) CheckPermissions(c *gin.Context) {
	user, _ := c.Get(global_const.CtxUserKey)

	switch c.FullPath() {
	case "/helpdesk/table/offices_list",
		"/helpdesk/table/update_schedule",
		"/helpdesk/table/schedule",
		"/helpdesk/table/lateness/",
		"/helpdesk/table/lateness/check_new",
		"/helpdesk/table/lateness/update":
		if !p.usecase.CheckPermission(user.(*models.User).Group.ID, actions.AdminTA) {
			c.AbortWithStatus(http.StatusForbidden)
		}
	}
}
