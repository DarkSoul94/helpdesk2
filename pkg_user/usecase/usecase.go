package usecase

import (
	"github.com/DarkSoul94/helpdesk2/global_const"
	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/DarkSoul94/helpdesk2/pkg/logger"
	"github.com/DarkSoul94/helpdesk2/pkg_user"
	"github.com/DarkSoul94/helpdesk2/pkg_user/perm_manager"
)

// NewUsecase ...
func NewUsecase(repo pkg_user.UserManagerRepo) *Usecase {
	perm, err := perm_manager.NewManager(global_const.Actions)
	if err != nil {
		logger.LogError("Init permissions manager", "user_manager/usecase", "", err)
	}
	uc := Usecase{
		repo:        repo,
		permManager: perm,
	}
	return &uc
}

func (u *Usecase) CreateUser(user *models.User) (uint64, error) {
	return u.repo.CreateUser(user)
}

func (u *Usecase) UserUpdate(author *models.User, userID, groupID uint64) error {
	//TODO add permissions check

	return u.repo.UserUpdate(userID, groupID)
}

func (u *Usecase) GetUserByEmail(email string) (models.User, error) {
	return u.repo.GetUserByEmail(email)
}

func (u *Usecase) GetUserByID(id uint64) (models.User, error) {
	return u.repo.GetUserByID(id)
}

func (u *Usecase) GetUsersList(user *models.User) ([]models.User, error) {
	var (
		err      error
		userList []models.User
	)

	//TODO add perm check

	if userList, err = u.repo.GetUsersList(); err != nil {
		return nil, ErrFailedGetUsersList
	}

	return userList, nil
}

func (u *Usecase) GetGroupByID(user *models.User, groupID uint64) (models.Group, error) {
	//TODO add permissions check
	return u.repo.GetGroupByID(groupID)
}

func (u *Usecase) GetGroupList(user *models.User) ([]models.Group, error) {
	//TODO add permissions check
	return u.repo.GetGroupList()
}

func (u *Usecase) GroupUpdate(id uint64, permission []byte) error {
	return nil
}

func (u *Usecase) CreateGroup(name string, permissions []byte) (uint64, error) {
	return 0, nil
}
