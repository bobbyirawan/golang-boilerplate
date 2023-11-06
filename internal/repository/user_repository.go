package repository

import (
	"go-chat/internal/model"

	"gorm.io/gorm"
)

type (
	UserRepository interface {
		GetUserByID(userID uint) (*model.User, error)
		GetUserByEmail(user *model.User) error
		CreateUser(user *model.User) error
		UpdateUser(user *model.User) error
		DeleteUser(userID uint) error
	}

	repository struct {
		db *gorm.DB
	}
)

func NewUserRepository(db *gorm.DB) UserRepository {
	return &repository{
		db: db,
	}
}

func (impl *repository) GetUserByID(userID uint) (*model.User, error) {
	user := &model.User{}
	if err := impl.db.First(user, userID).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (impl *repository) GetUserByEmail(user *model.User) error {
	if err := impl.db.Where("email = ?", user.Email).First(&user).Error; err != nil {
		return err
	}

	return nil
}

func (impl *repository) CreateUser(user *model.User) error {
	if err := impl.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (impl *repository) UpdateUser(user *model.User) error {
	if err := impl.db.Save(user).Error; err != nil {
		return err
	}
	return nil
}

func (impl *repository) DeleteUser(userID uint) error {
	if err := impl.db.Delete(&model.User{}, userID).Error; err != nil {
		return err
	}
	return nil
}
