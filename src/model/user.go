package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID            uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	LowestNoteId  int
	LowestNote    Note `json:"lowestNote" gorm:"foreignKey:LowestNoteId"`
	HighestNoteId int
	HighestNote   Note           `json:"highestNote" gorm:"foreignKey:HighestNoteId"`
	CreatedAt     time.Time      `json:"createdAt"`
	UpdatedAt     time.Time      `json:"updatedAt"`
	DeletedAt     gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}
