package http

import "github.com/DarkSoul94/helpdesk2/auth"

// Handler ...
type Handler struct {
	ucAuth auth.AuthUC
}

type inpGroup struct {
	ID          uint64                 `json:"group_id"`
	Name        string                 `json:"group_name"`
	Permissions map[string]interface{} `json:"permissions"`
}

type loginUser struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type inpUser struct {
	ID    uint64    `json:"user_id,omitempty"`
	Name  string    `json:"user_name,omitempty"`
	Token string    `json:"token"`
	Group *inpGroup `json:"group,omitempty"`
}
