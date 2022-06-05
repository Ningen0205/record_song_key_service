package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CoverSongKey struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	SongId    uuid.UUID
	Song      Song `json:"song" gorm:"unique;not null"`
	KeyId     int
	Key       Key            `json:"key"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}
