package model

import (
	"database/sql"
	"go-chat/pkg/utils"

	"gorm.io/gorm"
)

type User struct {
	ID          string         `gorm:"column:id;primaryKey;unique;<-:create"`
	Username    string         `gorm:"column:username;"`
	Description sql.NullString `gorm:"column:description"`
	Image       sql.NullString `gorm:"column:image"`
	Email       string         `gorm:"column:email;unique"`
	Password    string         `gorm:"column:password"`

	Default Default `gorm:"embedded"`
}

func (u *User) TableName() string {
	return "users"
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {

	user.ID = utils.GenerateUUID()

	return
}

// func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {

// 	return
// }
