package http

import (
	"net/http"

	"github.com/DarkSoul94/helpdesk2/global_const"
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
		"/helpdesk/support/active_support_list":
		if !p.usecase.CheckPermission(user.(*models.User).Group.ID, global_const.AdminTA) {
			c.AbortWithStatus(http.StatusForbidden)
		}
	case "/helpdesk/support/status_list",
		"/helpdesk/support/change_status",
		"/helpdesk/support/open_shift",
		"/helpdesk/support/close_shift":
		res1 := p.usecase.CheckPermission(user.(*models.User).Group.ID, global_const.AdminTA)
		res2 := p.usecase.CheckPermission(user.(*models.User).Group.ID, global_const.TicketTA_Work)
		if !res1 || !res2 {
			c.AbortWithStatus(http.StatusForbidden)
		}
	}
}
