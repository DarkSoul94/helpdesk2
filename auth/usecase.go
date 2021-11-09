package auth

import (
	"context"

	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/DarkSoul94/helpdesk2/pkg_user"
)

// Usecase ...
type AuthUC interface {
	LDAPSignIn(email, password string) (*pkg_user.User, string, models.Err)
	GenerateToken(user *pkg_user.User) (string, models.Err)
	ParseToken(ctx context.Context, accessToken string) (*pkg_user.User, models.Err)
}
