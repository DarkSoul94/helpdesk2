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

type OutUserWithOutGroup struct {
	ID         uint64 `json:"user_id"`
	Name       string `json:"user_name"`
	Email      string `json:"email"`
	Department string `json:"department,omitempty"`
	GroupID    uint64 `json:"group_id"`
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

func ToModelUser(user OutUser) *models.User {
	return &models.User{
		ID:         user.ID,
		Email:      user.Email,
		Name:       user.Name,
		Department: user.Department,
		Group:      ToModelGroup(user.Group),
	}
}

func ToOutUserWithOutGroup(user *models.User) OutUserWithOutGroup {
	return OutUserWithOutGroup{
		ID:         user.ID,
		Name:       user.Name,
		Email:      user.Email,
		Department: user.Department,
		GroupID:    user.Group.ID,
	}
}
