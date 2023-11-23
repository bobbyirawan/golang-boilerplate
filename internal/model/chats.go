package model

import (
	"go-chat/pkg/utils"
	"time"

	"gorm.io/gorm"
)

type Chat struct {
	ID        string    `gorm:"column:id;size:50;primaryKey;unique;<-:create"`
	Name      string    `gorm:"column:name;size:50;"`
	ChatType  string    `gorm:"column:chat_type;not null;default:1 on 1;"` // enum ('group', '1 on 1')
	CreatedAt time.Time `gorm:"column:created_at;not null;autoCreateTime;<-:create"`
}

func (c *Chat) TableName() string {
	return "chats"
}

func (c *Chat) BeforeCreate(tx *gorm.DB) (err error) {

	c.ID = utils.GenerateUUID()

	return
}
