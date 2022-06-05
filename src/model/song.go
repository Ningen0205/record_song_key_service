package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Song struct {
	ID             uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Name           string    `json:"name" gorm:"unique;not null"`
	NameKana       string    `json:"nameKana" gorm:"not null"`
	LowestNoteId   int
	LowestNote     Note `json:"lowestNote" gorm:"foreignKey:LowestNoteId"`
	HighestNoteId  int
	HighestNote    Note `json:"highestNote" gorm:"foreignKey:HighestNoteId"`
	SingerId       uuid.UUID
	Singer         Singer `json:"singer" gorm:"not null"`
	IsCover        bool   `json:"isCover" gorm:"not null"`
	OriginalSongId uuid.UUID
	Original       *Song          `json:"original" gorm:"foreignKey:OriginalSongId"`
	VideoLink      string         `json:"videoLink"`
	CreatedAt      time.Time      `json:"createdAt"`
	UpdatedAt      time.Time      `json:"updatedAt"`
	DeletedAt      gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}
