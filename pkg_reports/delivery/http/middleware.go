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
	case "/helpdesk/reports/motivation",
		"/helpdesk/reports/tickets_status_difference",
		"/helpdesk/reports/average_grades",
		"/helpdesk/reports/tickets_grades",
		"/helpdesk/reports/returned_tickets",
		"/helpdesk/reports/tickets_count",
		"/helpdesk/reports/supports_statuses":
		if !p.usecase.CheckPermission(user.(*models.User).Group.ID, actions.ReportTA_Get) {
			c.AbortWithStatus(http.StatusForbidden)
		}
	}
}
