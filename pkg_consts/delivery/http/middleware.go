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
	case "/helpdesk/const/:const":
		p.constsCheck(c, user.(*models.User))
	}
}

func (p *PermissionMiddleware) constsCheck(c *gin.Context, user *models.User) {
	var (
		constList []string = []string{
			"banner",
		}
		suppConst []string = []string{
			"banner",
		}
	)
	constant := c.Param("const")

	if !p.contains(constant, constList) {
		c.AbortWithStatus(http.StatusNotFound)
	}

	switch c.Request.Method {
	case "GET":
		return
	case "POST":
		if !p.usecase.CheckPermission(user.Group.ID, actions.AdminTA) && p.usecase.CheckPermission(user.Group.ID, actions.TicketTA_Work) {
			if !p.contains(constant, suppConst) {
				c.AbortWithStatus(http.StatusForbidden)
			}
		}
	default:
		c.AbortWithStatus(http.StatusMethodNotAllowed)
	}
}

func (p *PermissionMiddleware) contains(target string, slice []string) bool {
	for _, val := range slice {
		if val == target {
			return true
		}
	}
	return false
}
