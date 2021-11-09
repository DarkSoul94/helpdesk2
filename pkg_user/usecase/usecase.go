package usecase

import (
	"github.com/DarkSoul94/helpdesk2/pkg_user"
	"github.com/DarkSoul94/helpdesk2/pkg_user/group_manager"
)

// NewUsecase ...
func NewUsecase(repo pkg_user.UserManagerRepo, perm group_manager.PermManager) *Usecase {
	uc := Usecase{
		repo:        repo,
		permManager: perm,
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

	group, err := u.permManager.GetGroupByID(user.Group.ID)
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

	group, err := u.permManager.GetGroupByID(user.Group.ID)
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
		group, err := u.permManager.GetGroupByID(user.Group.ID)
		if err != nil {
			return nil, err
		}

		user.Group = group

		userList[id] = user
	}

	return userList, nil
}

func (u *Usecase) GetGroupByID(user *pkg_user.User, groupID uint64) (*group_manager.Group, error) {
	//TODO add permissions check
	return u.permManager.GetGroupByID(groupID)
}

func (u *Usecase) GetGroupList(user *pkg_user.User) ([]*group_manager.Group, error) {
	//TODO add permissions check
	return u.permManager.GetGroupList()
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
