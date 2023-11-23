package model

import (
	"go-chat/pkg/utils"

	"gorm.io/gorm"
)

type Contact struct {
	ID          string `gorm:"column:id;primaryKey;unique;<-:create"`
	UserID      string `gorm:"column:user_id"`
	RecipientID string `gorm:"column:recipient_id"`
	Username    string `gorm:"column:username;default:user"`
	Email       string `gorm:"column:email"`

	Default Default `gorm:"embedded"`
}

func (contact *Contact) TableName() string {
	return "contacts"
}

func (contact *Contact) BeforeCreate(tx *gorm.DB) (err error) {

	contact.ID = utils.GenerateUUID()

	return nil
}
