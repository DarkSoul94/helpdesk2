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
	case "/helpdesk/category/create",
		"/helpdesk/category/update",
		"/helpdesk/section/create",
		"/helpdesk/section/update",
		"/helpdesk/section/section_list":
		if !p.usecase.CheckPermission(user.(*models.User).Group.ID, global_const.AdminTA) {
			c.AbortWithStatus(http.StatusForbidden)
		}
	}
}
