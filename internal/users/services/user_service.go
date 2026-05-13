package services

import (
	"time"

	"github.com/marcelodiasdev/library-mvc/internal/users/models"
)

type UserService struct {
	userRepo models.UserRepository
}

func NewUserService(userRepo models.UserRepository) models.UserService {
	return &UserService{userRepo: userRepo}
}

func (u UserService) CreateUser(user *models.User) error {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	return u.userRepo.CreateUser(user)
}

func (u UserService) GetUser(id int64) (*models.User, error) {
	return u.userRepo.GetUser(id)
}

func (u UserService) GetAllUser() ([]*models.User, error) {
	return u.userRepo.GetAllUser()
}

func (u UserService) UpdatedUser(id int64, user *models.User) error {
	user.UpdatedAt = time.Now()
	return u.userRepo.UpdatedUser(id, user)
}

func (u UserService) DeleteUser(id int64) error {
	return u.userRepo.DeleteUser(id)
}
