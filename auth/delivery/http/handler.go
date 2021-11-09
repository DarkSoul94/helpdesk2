package http

import (
	"net/http"

	"github.com/DarkSoul94/helpdesk2/auth"
	"github.com/DarkSoul94/helpdesk2/dto"
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
		user loginUser
	)

	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"status": "error", "error": err.Error()})
		return
	}

	mUser, token, err := h.ucAuth.LDAPSignIn(user.UserName, user.Password)
	if err != nil {
		ctx.JSON(err.Code(), map[string]interface{}{"status": "error", "error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{"token": token, "user": dto.ToOutLoginUser(mUser, token)})
}
