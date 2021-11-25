package http

import (
	"net/http"

	"github.com/DarkSoul94/helpdesk2/dto"
	"github.com/DarkSoul94/helpdesk2/global_const"
	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/DarkSoul94/helpdesk2/pkg_user"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	ucUserManager pkg_user.IUserUsecase
}

// NewHandler ...
func NewHandler(uc pkg_user.IUserUsecase) *Handler {
	return &Handler{
		ucUserManager: uc,
	}
}

func (h *Handler) UpdateUser(ctx *gin.Context) {
	type updateUser struct {
		UserID  uint64 `json:"user_id"`
		GroupID uint64 `json:"group_id"`
	}
	var newUser updateUser

	user, _ := ctx.Get(global_const.CtxUserKey)

	if err := ctx.BindJSON(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"status": "error", "error": err.Error()})
		return
	}

	err := h.ucUserManager.UserUpdate(user.(*models.User), newUser.UserID, newUser.GroupID)
	if err != nil {
		ctx.JSON(err.Code(), map[string]interface{}{"status": "error", "error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{"status": "ok"})
}

func (h *Handler) GetUsersList(ctx *gin.Context) {
	var (
		userList []*models.User
		err      models.Err
	)

	user, _ := ctx.Get(global_const.CtxUserKey)

	if userList, err = h.ucUserManager.GetUsersList(user.(*models.User)); err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"status": "error", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.ToOutUserList(userList))
}

func (h *Handler) CreateGroup(ctx *gin.Context) {
	var group dto.OutGroup

	if err := ctx.BindJSON(&group); err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"status": "error", "error": err.Error()})
		return
	}

	user, _ := ctx.Get(global_const.CtxUserKey)
	id, err := h.ucUserManager.CreateGroup(user.(*models.User), dto.ToModelGroup(group))
	if err != nil {
		ctx.JSON(err.Code(), map[string]interface{}{"status": "error", "error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{"status": "ok", "group_id": id})
}

func (h *Handler) UpdateGroup(ctx *gin.Context) {
	var group dto.OutGroup

	if err := ctx.BindJSON(&group); err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"status": "error", "error": err.Error()})
		return
	}
	user, _ := ctx.Get(global_const.CtxUserKey)
	err := h.ucUserManager.GroupUpdate(user.(*models.User), dto.ToModelGroup(group))
	if err != nil {
		ctx.JSON(err.Code(), map[string]interface{}{"status": "error", "error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{"status": "ok"})
}

func (h *Handler) GetGroupsList(ctx *gin.Context) {
	user, _ := ctx.Get(global_const.CtxUserKey)

	groups, err := h.ucUserManager.GetGroupList(user.(*models.User))
	if err != nil {
		ctx.JSON(err.Code(), map[string]interface{}{"status": "error", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.ToOutGroupList(groups))
}

func (h *Handler) GetDepartmentsList(ctx *gin.Context) {
	departments, err := h.ucUserManager.GetDepartmentsList()
	if err != nil {
		ctx.JSON(err.Code(), map[string]string{"status": "error", "error": err.Error()})
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"status": "ok", "departments": departments})
}

func (h *Handler) GetGroupsListForResolve(c *gin.Context) {

}

/*
func (h *Handler) GetPermList(ctx *gin.Context) {
	perm, err := h.ucUserManager.GetFullPermListInBytes()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{Status: models.ErrStatus, Error: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, models.Response{Status: models.SuccessStatus, Data: h.toOutPermissions(perm)})
}
*/
