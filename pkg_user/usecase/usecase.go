package usecase

import (
	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/DarkSoul94/helpdesk2/pkg/logger"
	"github.com/DarkSoul94/helpdesk2/pkg_user"
	group_manager "github.com/DarkSoul94/helpdesk2/pkg_user/group_manager"
	groupUC "github.com/DarkSoul94/helpdesk2/pkg_user/group_manager/standart/usecase"
)

// NewUsecase ...
func NewUsecase(repo pkg_user.UserManagerRepo, grpRepo group_manager.GroupRepo) *Usecase {
	group, err := groupUC.NewManager(grpRepo)
	if err != nil {
		logger.LogError("Init permissions manager", "user_manager/usecase", "", err)
	}
	uc := Usecase{
		repo:  repo,
		group: group,
	}
	return &uc
}

func (u *Usecase) CreateUser(user *pkg_user.User) (uint64, error) {
	return u.repo.CreateUser(user)
}

func (u *Usecase) UserUpdate(author *pkg_user.User, userID, groupID uint64) error {
	//TODO add permissions check

	return u.repo.UpdateUser(userID, groupID)
}

func (u *Usecase) GetUserByEmail(email string) (*pkg_user.User, error) {
	user, err := u.repo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	group, err := u.group.GetGroupByID(user.Group.ID)
	if err != nil {
		return nil, err
	}

	user.Group = group

	return user, nil
}

func (u *Usecase) GetUserByID(id uint64) (*pkg_user.User, error) {
	user, err := u.repo.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	group, err := u.group.GetGroupByID(user.Group.ID)
	if err != nil {
		return nil, err
	}

	user.Group = group

	return user, nil
}

func (u *Usecase) GetUsersList(user *pkg_user.User) ([]*pkg_user.User, error) {
	var (
		err      error
		userList []*pkg_user.User
	)

	//TODO add perm check

	if userList, err = u.repo.GetUsersList(); err != nil {
		return nil, ErrFailedGetUsersList
	}

	for id, user := range userList {
		group, err := u.group.GetGroupByID(user.Group.ID)
		if err != nil {
			return nil, err
		}

		user.Group = group

		userList[id] = user
	}

	return userList, nil
}

func (u *Usecase) GetGroupByID(user *pkg_user.User, groupID uint64) (*group_manager.Group, models.Err) {
	//TODO add permissions check
	return u.group.GetGroupByID(groupID)
}

func (u *Usecase) GetGroupList(user *pkg_user.User) ([]*group_manager.Group, models.Err) {
	//TODO add permissions check
	return u.group.GetGroupList()
}

func (u *Usecase) GroupUpdate(id uint64, permission []byte) error {
	return nil
}

func (u *Usecase) CreateGroup(name string, permissions []byte) (uint64, error) {
	return 0, nil
}

func (u *Usecase) CheckPermissionForAction(user *pkg_user.User, actions ...string) bool {
	return true
}
