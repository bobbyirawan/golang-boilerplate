package model

import (
	"time"
)

type Default struct {
	CreatedAt time.Time  `gorm:"column:created_at;<-:create"`
	UpdatedAt *time.Time `gorm:"column:updated_at;autoUpdateTime;<-:update"`
}
