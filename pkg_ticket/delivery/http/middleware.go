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
	case "/helpdesk/category/create",
		"/helpdesk/category/update",
		"/helpdesk/section/create",
		"/helpdesk/section/update",
		"/helpdesk/section/section_list",
		"/helpdesk/region/create",
		"/helpdesk/region/update",
		"/helpdesk/region/",
		"/helpdesk/filial/create",
		"/helpdesk/filial/update",
		"/helpdesk/filial/",
		"/helpdesk/filial/filial_list":
		if !p.usecase.CheckPermission(user.(*models.User).Group.ID, actions.AdminTA) {
			c.AbortWithStatus(http.StatusForbidden)
		}

	case "/helpdesk/ticket/create":
		if !(p.usecase.CheckPermission(user.(*models.User).Group.ID, actions.AdminTA) || p.usecase.CheckPermission(user.(*models.User).Group.ID, actions.TicketTA_Work)) {
			c.AbortWithStatus(http.StatusForbidden)
		}

	case "/helpdesk/resolve_ticket/check_exist",
		"/helpdesk/resolve_ticket/resolve_tickets_list":
		if !p.usecase.CheckPermission(user.(*models.User).Group.ID, actions.TicketTA_Resolve) && !p.usecase.CheckPermission(user.(*models.User).Group.ID, actions.TicketTA_Work) {
			c.AbortWithStatus(http.StatusForbidden)
		}
	}
}
