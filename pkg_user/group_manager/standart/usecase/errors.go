package usecase

import "github.com/DarkSoul94/helpdesk2/models"

var (
	GroupErr_Exist = models.BadRequest("Group already exist")
)
