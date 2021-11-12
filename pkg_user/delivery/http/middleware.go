package http

import (
	"net/http"

	"github.com/DarkSoul94/helpdesk2/global_const"
	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/DarkSoul94/helpdesk2/pkg_user"
	"github.com/gin-gonic/gin"
)

type PermissionMiddleware struct {
	usecase pkg_user.IUserUsecase
}

func NewPermissionMiddleware(usecase pkg_user.IUserUsecase) gin.HandlerFunc {
	return (&PermissionMiddleware{
		usecase: usecase,
	}).CheckPermissions
}

func (p *PermissionMiddleware) CheckPermissions(c *gin.Context) {
	user, _ := c.Get(global_const.CtxUserKey)

	switch c.FullPath() {
	case "/helpdesk/group/create",
		"/helpdesk/group/update",
		"/helpdesk/group/",
		"/helpdesk/category/create",
		"/helpdesk/category/update":
		if err := p.usecase.CheckPermissions(user.(*models.User), global_const.AdminTA); err != nil {
			c.AbortWithStatus(http.StatusForbidden)
		}

	case "/helpdesk/user/":
		err1 := p.usecase.CheckPermissions(user.(*models.User), global_const.AdminTA)
		err2 := p.usecase.CheckPermissions(user.(*models.User), global_const.TicketTA_FullSearch)
		if err1 != nil || err2 != nil {
			c.AbortWithStatus(http.StatusForbidden)
		}
	}
}
