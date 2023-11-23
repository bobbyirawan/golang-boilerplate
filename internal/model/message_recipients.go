package model

import (
	"go-chat/pkg/utils"

	"gorm.io/gorm"
)

var (
	MESSAGE_RECIPIENT_SEND     = "terkirim"
	MESSAGE_RECIPIENT_NOT_SEND = "belum terkirim"
	MESSAGE_RECIPIENT_READ     = "dibaca"
)

type MessageRecipient struct {
	ID          string `gorm:"column:id;size:50;primaryKey;unique;<-:create"`
	MessageID   string `gorm:"column:message_id;size:50;not null"`
	RecipientID string `gorm:"column:recipient_id;size:50;not null"`
	Status      string `gorm:"column:status;default:belum terkirim"` //enum('belum terkirim','terkirim','dibaca')

	Default Default `gorm:"embedded"`
}

func (*MessageRecipient) TableName() string {
	return "message_recipients"
}

func (mr *MessageRecipient) BeforeCreate(tx *gorm.DB) (err error) {

	mr.ID = utils.GenerateUUID()

	return
}
