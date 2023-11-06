package model

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uint64    `gorm:"primaryKey" json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email" gorm:"not null;unique"`
	Password  string    `json:"-"`
	CreatedAt time.Time `gorm:"->:false;<-:create" json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}
