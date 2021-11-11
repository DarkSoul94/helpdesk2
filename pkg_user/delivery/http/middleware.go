package http

import (
	"net/http"

	"github.com/DarkSoul94/helpdesk2/global_const"
	"github.com/DarkSoul94/helpdesk2/pkg_user"
	"github.com/gin-gonic/gin"
)

type PermissionMiddleware struct {
	usecase pkg_user.UserManagerUC
}

func NewPermissionMiddleware(usecase pkg_user.UserManagerUC) gin.HandlerFunc {
	return (&PermissionMiddleware{
		usecase: usecase,
	}).CheckPermissions
}

func (p *PermissionMiddleware) CheckPermissions(c *gin.Context) {
	user, _ := c.Get(global_const.CtxUserKey)

	switch c.FullPath() {
	case "/helpdesk/group/create",
		"/helpdesk/group/update",
		"/helpdesk/group/":
		if err := p.usecase.CheckPermissionForAction(user.(*pkg_user.User), global_const.AdminTA); err != nil {
			c.AbortWithStatus(http.StatusForbidden)
		}

	case "/helpdesk/user/":
		err1 := p.usecase.CheckPermissionForAction(user.(*pkg_user.User), global_const.AdminTA)
		err2 := p.usecase.CheckPermissionForAction(user.(*pkg_user.User), global_const.TicketTA_FullSearch)
		if err1 != nil || err2 != nil {
			c.AbortWithStatus(http.StatusForbidden)
		}
	}
}
