package mysql

import "github.com/DarkSoul94/helpdesk2/models"

var (
	GroupErr_GetList = models.InternalError("Failed to get groups list")
	GroupErr_Create  = models.InternalError("Failed to create new group")

	GroupErr_Exist    = models.BadRequest("Group already exist")
	GroupErr_NotFound = models.BadRequest("Group not found")
)
