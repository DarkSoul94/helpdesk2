package http

import (
	"net/http"

	"github.com/DarkSoul94/helpdesk2/auth"
	"github.com/DarkSoul94/helpdesk2/dto"
	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/DarkSoul94/helpdesk2/pkg_user"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	ucAuth auth.AuthUC
}

// NewHandler ...
func NewHandler(ucAuth auth.AuthUC) *Handler {
	return &Handler{
		ucAuth: ucAuth,
	}
}

type loginUser struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

//SignIn ...
func (h *Handler) SignIn(ctx *gin.Context) {
	var (
		user  loginUser
		mUser *pkg_user.User
		token string
		err   error
	)

	err = ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{Status: models.ErrStatus, Error: err.Error()})
		return
	}

	mUser, token, err = h.ucAuth.LDAPSignIn(user.UserName, user.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{Status: models.ErrStatus, Error: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, models.Response{Status: models.ErrStatus, Data: dto.ToOutLoginUser(mUser, token)})
}
