package pkg_user

type UserManagerRepo interface {
	CreateUser(user *User) (uint64, error)
	UpdateUser(userID, groupID uint64) error
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id uint64) (*User, error)
	GetUsersList() ([]*User, error)

	Close() error
}
