package dto

import "github.com/DarkSoul94/helpdesk2/pkg_user"

type OutUser struct {
	ID         uint64   `json:"id"`
	Name       string   `json:"name"`
	Email      string   `json:"email"`
	Department string   `json:"department"`
	Group      OutGroup `json:"group"`
}

type OutLoginUser struct {
	ID         uint64   `json:"user_id"`
	Name       string   `json:"user_name"`
	Email      string   `json:"email"`
	Department string   `json:"department"`
	Group      OutGroup `json:"group"`
	Token      string   `json:"token"`
}

type OutUserForList struct {
	ID         uint64 `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Department string `json:"department"`
	GroupID    uint64 `json:"group_id"`
}

func ToOutUser(user *pkg_user.User) OutUser {
	return OutUser{
		ID:         user.ID,
		Name:       user.Name,
		Email:      user.Email,
		Department: user.Department,
		Group:      ToOutGroup(user.Group),
	}
}

func ToOutLoginUser(user *pkg_user.User, token string) OutLoginUser {
	return OutLoginUser{
		ID:         user.ID,
		Name:       user.Name,
		Email:      user.Email,
		Department: user.Department,
		Group:      ToOutGroup(user.Group),
		Token:      token,
	}
}

func ToOutUserList(users []*pkg_user.User) []OutUserForList {
	var outUsers []OutUserForList = make([]OutUserForList, 0)

	for _, user := range users {
		outUsers = append(outUsers, toOutUserForList(user))
	}

	return outUsers
}

func toOutUserForList(user *pkg_user.User) OutUserForList {
	return OutUserForList{
		ID:         user.ID,
		Name:       user.Name,
		Email:      user.Email,
		Department: user.Department,
		GroupID:    user.Group.ID,
	}
}
