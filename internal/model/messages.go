package model

import (
	"go-chat/pkg/utils"

	"gorm.io/gorm"
)

type Message struct {
	ID       string `gorm:"column:id;size:50;primaryKey;unique;<-:create"`
	ChatID   string `gorm:"column:chat_id;size:50;not null"`
	SenderID string `gorm:"column:sender_id;size:50;not null"`
	Content  string `gorm:"column:content"`
	Status   string `gorm:"column:status;default:belum terkirim"` //enum('belum terkirim','terkirim','dibaca')

	Default Default `gorm:"embedded"`
}

func (*Message) TableName() string {
	return "messages"
}

func (message *Message) BeforeCreate(tx *gorm.DB) (err error) {

	message.ID = utils.GenerateUUID()

	return
}
