package repository

import (
	"go-chat/internal/model"

	"gorm.io/gorm"
)

type (
	MessageRepository interface {
		Create(message *model.Message) error
		UpdateStatusSend(ID string) error
	}

	messageDependency struct {
		mysql *gorm.DB
	}
)

func NewMessageRepository(mysql *gorm.DB) MessageRepository {
	return &messageDependency{
		mysql: mysql,
	}
}

func (impl *messageDependency) Create(message *model.Message) error {
	return impl.mysql.Create(message).Error
}

func (impl *messageDependency) UpdateStatusSend(ID string) error {
	message := &model.Message{}
	message.ID = ID
	return impl.mysql.Model(message).Update("status", "terkirim").Error
}
