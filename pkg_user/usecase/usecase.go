package usecase

import (
	"github.com/DarkSoul94/helpdesk2/global_const"
	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/DarkSoul94/helpdesk2/pkg_support"
	"github.com/DarkSoul94/helpdesk2/pkg_user"
	group_manager "github.com/DarkSoul94/helpdesk2/pkg_user/group_manager"
)

// NewUsecase ...
func NewUsecase(
	repo pkg_user.IUserRepo,
	group group_manager.IGroupUsecase,
	perm group_manager.IPermManager,
	support pkg_support.ISupportUsecase,
) *Usecase {
	return &Usecase{
		repo:    repo,
		group:   group,
		perm:    perm,
		support: support,
	}
}

func (u *Usecase) CreateUser(user *models.User) (uint64, models.Err) {
	return u.repo.CreateUser(user)
}

func (u *Usecase) UserUpdate(askUser *models.User, userID, groupID uint64) models.Err {
	//проверка наличия доступа у запрашивающего на изменение целевого пользователя
	if !u.perm.CheckPermission(askUser.Group.ID, global_const.AdminTA_UserUpdate) {
		return errPermissions_UserUpdate
	}
	tUser, err := u.repo.GetUserByID(userID)
	if err != nil {
		return err
	}
	//Проверка наличия прав на работу с запросами у текущей (forCurrent) и новой (forNew) групп
	forCurrent := u.perm.CheckPermission(tUser.Group.ID, global_const.TicketTA_Work)
	forNew := u.perm.CheckPermission(groupID, global_const.TicketTA_Work)

	//если у текущей группы целевого пользователя нет прав на обработку запросов в ТП,
	//а у новой группы есть, то добавляем пользователя в саппорты
	if !forCurrent && forNew {
		if err := u.support.CreateSupport(userID); err != nil {
			return err
		}
	}

	//если же у текущей группы есть права на обработку запроса,
	//а у новой - нет, то удаляем пользователя из саппортов
	if forCurrent && !forNew {
		if err := u.support.DeleteSupport(userID); err != nil {
			return err
		}
	}
	return u.repo.UpdateUser(userID, groupID)
}

func (u *Usecase) fillGroup(user *models.User) models.Err {
	group, err := u.group.GetGroupByID(user.Group.ID)
	if err != nil {
		return err
	}
	user.Group = group

	return nil
}

func (u *Usecase) GetUserByEmail(email string) (*models.User, models.Err) {
	user, err := u.repo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	if err := u.fillGroup(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *Usecase) GetUserByID(id uint64) (*models.User, models.Err) {
	user, err := u.repo.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	if err := u.fillGroup(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *Usecase) GetUsersList(askUser *models.User) ([]*models.User, models.Err) {
	var (
		err      models.Err
		userList []*models.User
	)

	if !u.perm.CheckPermission(askUser.Group.ID, global_const.AdminTA_UserGet) {
		return nil, err
	}

	if userList, err = u.repo.GetUsersList(); err != nil {
		return nil, err
	}
	for _, user := range userList {
		if err := u.fillGroup(user); err != nil {
			return nil, err
		}
	}

	return userList, nil
}

func (u *Usecase) GetGroupList(askUser *models.User) ([]*models.Group, models.Err) {
	if !u.perm.CheckPermission(askUser.Group.ID, global_const.AdminTA_GroupGet) {
		return nil, errPermissions_GetGroupList
	}
	return u.group.GetGroupList()
}

func (u *Usecase) GroupUpdate(askUser *models.User, group *models.Group) models.Err {
	if !u.perm.CheckPermission(askUser.Group.ID, global_const.AdminTA_GroupUpdate) {
		return errPermissions_UpdateGroup
	}
	forCurrent := u.perm.CheckPermission(group.ID, global_const.TicketTA_Work)
	forNew := u.perm.CheckUpdatedPermissions(group, global_const.TicketTA_Work)
	//если текущая группа не обладает таким правом, проводится проверка обновленной группы на его наличие
	//и если такое право есть то выбираются ID пользователей входящих в эту группу и они добавляются в саппорты
	if !forCurrent && forNew {
		users, err := u.group.GetUsersByGroup(group.ID)
		if err != nil {
			return err
		}
		if err = u.support.CreateSupport(users...); err != nil {
			return err
		}
	}

	//если текущая группа обладает таким правом, проводится проверка обновленной группы на его отсутствие
	//и если такое право отсутствует то выбираются ID пользователей входящих в эту группу и они исключаются из саппортов
	if forCurrent && !forNew {
		users, err := u.group.GetUsersByGroup(group.ID)
		if err != nil {
			return err
		}
		if err = u.support.DeleteSupport(users...); err != nil {
			return err
		}
	}
	return u.group.GroupUpdate(group)
}

func (u *Usecase) CreateGroup(askUser *models.User, group *models.Group) (uint64, models.Err) {
	if !u.perm.CheckPermission(askUser.Group.ID, global_const.AdminTA_GroupCreate) {
		return 0, errPermissions_CreateGroup
	}
	return u.group.CreateGroup(group)
}

func (u *Usecase) GetDepartmentsList() ([]string, models.Err) {
	return u.repo.GetDepartmentsList()
}
