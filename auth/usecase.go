package auth

import (
	"context"

	"github.com/DarkSoul94/helpdesk2/models"
)

// Usecase ...
type AuthUC interface {
	LDAPSignIn(email, password string) (*models.User, string, models.Err)
	GenerateToken(user *models.User) (string, models.Err)
	ParseToken(ctx context.Context, accessToken string) (*models.User, models.Err)
}
