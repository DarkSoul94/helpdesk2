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
	case "/helpdesk/support/support_list",
		"/helpdesk/support/active_support_list",
		"/helpdesk/support/card/update",
		"/helpdesk/support/card/seniors",
		"/helpdesk/support/card/list",
		"/helpdesk/support/card":
		if !p.usecase.CheckPermission(user.(*models.User).Group.ID, actions.AdminTA) {
			c.AbortWithStatus(http.StatusForbidden)
		}
	case "/helpdesk/support/status_list",
		"/helpdesk/support/change_status",
		"/helpdesk/support/open_shift",
		"/helpdesk/support/create_lateness",
		"/helpdesk/support/close_shift",
		"/helpdesk/support/shift_status",
		"/helpdesk/support/get_current_statuses",
		"/helpdesk/support/get_support_status":
		res1 := p.usecase.CheckPermission(user.(*models.User).Group.ID, actions.AdminTA)
		res2 := p.usecase.CheckPermission(user.(*models.User).Group.ID, actions.TicketTA_Work)
		if !(res1 || res2) {
			c.AbortWithStatus(http.StatusForbidden)
		}
	}
}
