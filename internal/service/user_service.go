package service

import (
	"go-boilerplate/internal/repository"
	"go-boilerplate/model"
)

type (
	UserService interface {
		GetUser(userID uint) (*model.User, error)
		CreateUser(user *model.User) error
		UpdateUser(user *model.User) error
		DeleteUser(userID uint) error
	}

	userService struct {
		repo repository.Holder
	}
)

func NewUserService(repo repository.Holder) UserService {
	return &userService{
		repo: repo,
	}
}

func (impl *userService) GetUser(userID uint) (*model.User, error) {
	return impl.repo.UserRepository.GetUserByID(userID)
}

func (impl *userService) CreateUser(user *model.User) error {
	return impl.repo.UserRepository.CreateUser(user)
}

func (impl *userService) UpdateUser(user *model.User) error {
	return impl.repo.UserRepository.UpdateUser(user)
}

func (impl *userService) DeleteUser(userID uint) error {
	return impl.repo.UserRepository.DeleteUser(userID)
}
