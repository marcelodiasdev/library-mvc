package models

type UserService interface {
	CreateUser(user *User) error
	GetUser(id int64) (*User, error)
	GetAllUser() ([]*User, error)
	UpdatedUser(int int64, user *User) error
	DeleteUser(int int64) error
}
