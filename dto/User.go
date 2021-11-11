package dto

import (
	"github.com/DarkSoul94/helpdesk2/models"
)

type OutUser struct {
	ID         uint64   `json:"user_id"`
	Name       string   `json:"user_name"`
	Email      string   `json:"email"`
	Department string   `json:"department,omitempty"`
	Group      OutGroup `json:"group"`
}

func ToOutUser(user *models.User) OutUser {
	return OutUser{
		ID:         user.ID,
		Name:       user.Name,
		Email:      user.Email,
		Department: user.Department,
		Group:      ToOutGroup(user.Group),
	}
}

func ToOutLoginUser(user *models.User) OutUser {
	return OutUser{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Group: ToOutGroup(user.Group),
	}
}

func ToOutUserList(users []*models.User) []OutUser {
	var outUsers []OutUser = make([]OutUser, 0)

	for _, user := range users {
		outUsers = append(outUsers, ToOutUser(user))
	}

	return outUsers
}
