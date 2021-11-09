package http

import (
	"net/http"
	"strconv"

	"github.com/DarkSoul94/helpdesk2/dto"
	"github.com/DarkSoul94/helpdesk2/global_const"
	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/DarkSoul94/helpdesk2/pkg_user"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	ucUserManager pkg_user.UserManagerUC
}

// NewHandler ...
func NewHandler(uc pkg_user.UserManagerUC) *Handler {
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
		ctx.JSON(http.StatusBadRequest, models.Response{Status: models.ErrStatus, Error: err.Error()})
		return
	}

	err = h.ucUserManager.UserUpdate(user.(*pkg_user.User), newUser.UserID, newUser.GroupID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{Status: models.ErrStatus, Error: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, models.Response{Status: models.SuccessStatus})
}

func (h *Handler) GetUsersList(ctx *gin.Context) {
	var (
		userList []*pkg_user.User
		err      error
	)

	user, _ := ctx.Get(global_const.CtxUserKey)

	if userList, err = h.ucUserManager.GetUsersList(user.(*pkg_user.User)); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{Status: models.ErrStatus, Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{Status: models.SuccessStatus, Data: dto.ToOutUserList(userList)})
}

func (h *Handler) CreateGroup(ctx *gin.Context) {

}

func (h *Handler) UpdateGroup(ctx *gin.Context) {

}

func (h *Handler) GetGroup(ctx *gin.Context) {
	groupID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{Status: models.ErrStatus, Error: err.Error()})
		return
	}

	user, _ := ctx.Get(global_const.CtxUserKey)

	group, err := h.ucUserManager.GetGroupByID(user.(*pkg_user.User), groupID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{Status: models.ErrStatus, Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{Status: models.SuccessStatus, Data: dto.ToOutGroup(group)})
}

func (h *Handler) GetGroupsList(ctx *gin.Context) {
	user, _ := ctx.Get(global_const.CtxUserKey)

	groups, err := h.ucUserManager.GetGroupList(user.(*pkg_user.User))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{Status: models.ErrStatus, Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{Status: models.SuccessStatus, Data: dto.ToOutGroupList(groups)})
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
