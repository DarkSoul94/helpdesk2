package http

import "github.com/DarkSoul94/helpdesk2/pkg_user"

// Handler ...
type Handler struct {
	ucUserManager pkg_user.UserManagerUC
}

type Response struct {
	Status string      `json:"status"`
	Error  string      `json:"error,omitempty"`
	Data   interface{} `json:"data"`
}

type outUser struct {
	ID    uint64    `json:"id"`
	Name  string    `json:"name"`
	Group *outGroup `json:"group"`
}

type outUserForList struct {
	ID         uint64 `json:"id"`
	Name       string `json:"name"`
	GroupID    uint64 `json:"group_id"`
	Department string `json:"department"`
}

type outGroup struct {
	ID          uint64                 `json:"id"`
	Name        string                 `json:"name"`
	Permissions map[string]interface{} `json:"permissions"`
}
