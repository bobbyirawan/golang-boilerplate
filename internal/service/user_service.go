package service

import (
	"go-chat/internal/dto"
	"go-chat/internal/model"
	"go-chat/internal/repository"
)

type (
	UserService interface {
		GetUser(req *dto.GetUserByIDReq, res *dto.GetUserByIDRes) error
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

func (impl *userService) GetUser(req *dto.GetUserByIDReq, res *dto.GetUserByIDRes) error {
	user := new(model.User)
	user.ID = req.UserID

	if err := impl.repo.UserRepository.GetUserByID(user); err != nil {
		return err
	}

	res.Description = user.Description.String
	res.Email = user.Email
	res.Image = user.Image.String
	res.Username = user.Username

	return nil
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
