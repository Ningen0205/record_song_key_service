package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserSong struct {
	UserId    uuid.UUID
	User      User `json:"user" gorm:"not null;primarykey"`
	SongId    uuid.UUID
	Song      Song `json:"song" gorm:"not null;primarykey"`
	KeyId     int
	Key       Key            `json:"key" `
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}
