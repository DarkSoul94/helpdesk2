package http

import (
	"net/http"

	"github.com/DarkSoul94/helpdesk2/auth"
	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/gin-gonic/gin"
)

// NewHandler ...
func NewHandler(ucAuth auth.AuthUC) *Handler {
	return &Handler{
		ucAuth: ucAuth,
	}
}

//SignIn ...
func (h *Handler) SignIn(ctx *gin.Context) {
	var (
		user    loginUser
		mUser   models.User
		outUser inpUser
		token   string
		err     error
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
	outUser = h.toInpUser(mUser, token)
	ctx.JSON(http.StatusOK, models.Response{Status: models.ErrStatus, Data: outUser})
}
