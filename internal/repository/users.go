package repository

import (
	"go-chat/internal/model"

	"gorm.io/gorm"
)

type (
	UserRepository interface {
		GetUserByID(user *model.User) error
		GetUserByEmail(user *model.User) error
		CreateUser(user *model.User) error
		UpdateUser(user *model.User) error
		DeleteUser(userID uint) error
	}

	userRepository struct {
		db *gorm.DB
	}
)

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (impl *userRepository) GetUserByID(user *model.User) error {
	return impl.db.First(user, "id = ? ", user.ID).Error
}

func (impl *userRepository) GetUserByEmail(user *model.User) error {
	if err := impl.db.Where("email = ?", user.Email).First(&user).Error; err != nil {
		return err
	}
	return nil
}

func (impl *userRepository) CreateUser(user *model.User) error {
	if err := impl.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (impl *userRepository) UpdateUser(user *model.User) error {
	if err := impl.db.Save(user).Error; err != nil {
		return err
	}
	return nil
}

func (impl *userRepository) DeleteUser(userID uint) error {
	if err := impl.db.Delete(&model.User{}, userID).Error; err != nil {
		return err
	}
	return nil
}
