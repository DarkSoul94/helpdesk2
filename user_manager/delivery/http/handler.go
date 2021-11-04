package http

import (
	"net/http"
	"strconv"

	"github.com/DarkSoul94/helpdesk2/global_const"
	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/DarkSoul94/helpdesk2/user_manager"
	"github.com/gin-gonic/gin"
)

// NewHandler ...
func NewHandler(uc user_manager.UserManagerUC) *Handler {
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
	err := ctx.BindJSON(&newUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, Response{Status: models.ErrStatus, Error: err.Error()})
		return
	}

	err = h.ucUserManager.UserUpdate(user.(*models.User), newUser.UserID, newUser.GroupID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, Response{Status: models.ErrStatus, Error: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, Response{Status: models.SuccessStatus})
}

func (h *Handler) GetUsersList(ctx *gin.Context) {
	var (
		outUsers []outUserForList
		userList []models.User
		err      error
	)

	user, _ := ctx.Get(global_const.CtxUserKey)
	if userList, err = h.ucUserManager.GetUsersList(user.(*models.User)); err != nil {
		ctx.JSON(http.StatusBadRequest, Response{Status: models.ErrStatus, Error: err.Error()})
		return
	}
	for _, user := range userList {
		outUsers = append(outUsers, h.toOutUserForList(user))
	}

	ctx.JSON(http.StatusOK, Response{Status: models.SuccessStatus, Data: outUsers})
}

func (h *Handler) CreateGroup(ctx *gin.Context) {

}

func (h *Handler) UpdateGroup(ctx *gin.Context) {

}

func (h *Handler) GetGroup(ctx *gin.Context) {
	id := ctx.Param("id")
	groupID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, Response{Status: models.ErrStatus, Error: err.Error()})
		return
	}
	user, _ := ctx.Get(global_const.CtxUserKey)

	group, err := h.ucUserManager.GetGroupByID(user.(*models.User), uint64(groupID))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, Response{Status: models.ErrStatus, Error: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, Response{Status: models.SuccessStatus, Data: h.toOutGroup(group)})
}

func (h *Handler) GetGroupsList(ctx *gin.Context) {
	var outGroups []outGroup
	user, _ := ctx.Get(global_const.CtxUserKey)
	groups, err := h.ucUserManager.GetGroupList(user.(*models.User))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, Response{Status: models.ErrStatus, Error: err.Error()})
		return
	}
	for _, val := range groups {
		outGroups = append(outGroups, h.toOutGroupForList(val))
	}
	ctx.JSON(http.StatusOK, Response{Status: models.SuccessStatus, Data: outGroups})

}

func (h *Handler) GetPermList(ctx *gin.Context) {
	perm, err := h.ucUserManager.GetFullPermListInBytes()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, Response{Status: models.ErrStatus, Error: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, Response{Status: models.SuccessStatus, Data: h.toOutPermissions(perm)})
}
