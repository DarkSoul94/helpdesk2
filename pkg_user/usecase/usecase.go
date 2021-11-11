package usecase

import (
	"github.com/DarkSoul94/helpdesk2/global_const"
	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/DarkSoul94/helpdesk2/pkg/logger"
	"github.com/DarkSoul94/helpdesk2/pkg_user"
	group_manager "github.com/DarkSoul94/helpdesk2/pkg_user/group_manager"
	groupUC "github.com/DarkSoul94/helpdesk2/pkg_user/group_manager/standart/usecase"
	"github.com/DarkSoul94/helpdesk2/pkg_user/support_manager"
	suppUC "github.com/DarkSoul94/helpdesk2/pkg_user/support_manager/usecase"
)

// NewUsecase ...
func NewUsecase(
	repo pkg_user.UserManagerRepo,
	grpRepo group_manager.GroupRepo,
	suppRepo support_manager.SupportRepo,
) *Usecase {

	group, err := groupUC.NewManager(grpRepo)
	if err != nil {
		logger.LogError("Init group manager", "user_manager/usecase", "", err)
	}
	support := suppUC.NewSupportUsecase(suppRepo)
	uc := Usecase{
		repo:    repo,
		group:   group,
		support: support,
	}
	return &uc
}

func (u *Usecase) CreateUser(user *models.User) (uint64, models.Err) {
	return u.repo.CreateUser(user)
}

func (u *Usecase) UserUpdate(askUser *models.User, userID, groupID uint64) models.Err {
	//TODO add permissions check

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

	//TODO add perm check

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
	//TODO add permissions check
	return u.group.GetGroupList()
}

func (u *Usecase) GroupUpdate(id uint64, permission []byte) models.Err {
	return nil
}

func (u *Usecase) CreateGroup(askUser *models.User, group *models.Group) (uint64, models.Err) {
	if err := u.group.CheckPermission(askUser.Group.ID, global_const.AdminTA_GroupCreate); err != nil {
		return 0, err
	}
	return u.group.CreateGroup(group)
}

func (u *Usecase) CheckPermissions(user *models.User, actions ...string) models.Err {
	return nil
}
