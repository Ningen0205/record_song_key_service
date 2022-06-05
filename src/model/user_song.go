package model

import (
	"time"

	"gorm.io/gorm"
)

type UserSong struct {
	User      User           `json:"user" gorm:"not null;primarykey"`
	Song      Song           `json:"song" gorm:"not null;primarykey"`
	Key       Key            `json:"key"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}
