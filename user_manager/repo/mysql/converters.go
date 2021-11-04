package mysql

import (
	"encoding/json"
	"strconv"

	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/DarkSoul94/helpdesk2/pkg/logger"
	"github.com/DarkSoul94/helpdesk2/user_manager/perm_manager"
)

func (r *Repo) toModelGroup(dbGroup dbGroup) models.Group {
	var temp perm_manager.PermLayer
	err := json.Unmarshal(dbGroup.Permissions, &temp)
	if err != nil {
		logger.LogError(ErrReadGroup.Error(), "user_manager/repo/mysql", strconv.FormatUint(dbGroup.ID, 10), err)
	}
	mGroup := models.Group{
		ID:          dbGroup.ID,
		Name:        dbGroup.Name,
		Permissions: temp,
	}
	return mGroup
}

func (r *Repo) toModelUser(dbUser *dbUser) models.User {
	var group models.Group
	if dbUser.GroupID != 0 {
		group = models.Group{ID: dbUser.GroupID}
	} else {
		group = models.Group{}
	}
	return models.User{
		ID:         dbUser.ID,
		Email:      dbUser.Email,
		Name:       dbUser.Name,
		Group:      &group,
		Department: dbUser.Department,
	}
}

func (r *Repo) toDBUser(mUser *models.User) dbUser {
	user := dbUser{
		ID:         mUser.ID,
		Email:      mUser.Email,
		Name:       mUser.Name,
		Department: mUser.Department,
	}
	if mUser.Group != nil {
		user.GroupID = mUser.Group.ID
	}
	return user
}
