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
	if err := u.perm.CheckPermission(askUser.Group.ID, global_const.AdminTA_UserUpdate); err != nil {
		return err
	}
	tUser, err := u.repo.GetUserByID(userID)
	if err != nil {
		return err
	}
	//если у текущей группы целевого пользователя нет прав на обработку запросов в ТП,
	//а у новой группы есть, то добавляем пользователя в саппорты
	if err := u.perm.CheckPermission(tUser.Group.ID, global_const.TicketTA_Work); err != nil {
		if err := u.perm.CheckPermission(groupID, global_const.TicketTA_Work); err == nil {
			if err := u.support.CreateSupport(userID); err != nil {
				return err
			}
		}
	} else {
		//если же у текущей группы есть права на обработку запроса, а у новой - нет, то удаляем пользователя из саппортов
		if err := u.perm.CheckPermission(groupID, global_const.TicketTA_Work); err != nil {
			if err := u.support.DeleteSupport(userID); err != nil {
				return err
			}
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

	if err := u.perm.CheckPermission(askUser.Group.ID, global_const.AdminTA_UserGet); err != nil {
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
	if err := u.perm.CheckPermission(askUser.Group.ID, global_const.AdminTA_GroupGet); err != nil {
		return nil, err
	}
	return u.group.GetGroupList()
}

func (u *Usecase) GroupUpdate(askUser *models.User, group *models.Group) models.Err {
	if err := u.perm.CheckPermission(askUser.Group.ID, global_const.AdminTA_GroupUpdate); err != nil {
		return err
	}
	//проверка что существующая группа имеет права на обработку запроса
	if err := u.perm.CheckPermission(group.ID, global_const.TicketTA_Work); err != nil {
		//если текущая группа не обладает таким правом, проводится проверка обновленной группы на его наличие
		//и если такое право есть то выбираются ID пользователей входящих в эту группу и они добавляются в саппорты
		if err := u.perm.CheckUpdatedPermissions(group, global_const.TicketTA_Work); err == nil {
			if users, err := u.group.GetUsersByGroup(group.ID); err != nil {
				return err
			} else {
				if err = u.support.CreateSupport(users...); err != nil {
					return err
				}
			}
		}
	} else {
		//если текущая группа обладает таким правом, проводится проверка обновленной группы на его отсутствие
		//и если такое право отсутствует то выбираются ID пользователей входящих в эту группу и они исключаются из саппортов
		if err := u.perm.CheckUpdatedPermissions(group, global_const.TicketTA_Work); err != nil {
			if users, err := u.group.GetUsersByGroup(group.ID); err != nil {
				return err
			} else {
				if err = u.support.DeleteSupport(users...); err != nil {
					return err
				}
			}
		}
	}

	return u.group.GroupUpdate(group)
}

func (u *Usecase) CreateGroup(askUser *models.User, group *models.Group) (uint64, models.Err) {
	if err := u.perm.CheckPermission(askUser.Group.ID, global_const.AdminTA_GroupCreate); err != nil {
		return 0, err
	}
	return u.group.CreateGroup(group)
}
