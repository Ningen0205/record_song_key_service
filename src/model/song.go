package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Song struct {
	ID          uuid.UUID      `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Name        string         `json:"name" gorm:"unique;not null"`
	NameKana    string         `json:"nameKana" gorm:"not null"`
	LowestNote  Note           `json:"lowestNote"`
	HighestNote Note           `json:"highestNote"`
	Singer      Singer         `json:"singer" gorm:"not null"`
	IsCover     bool           `json:"isCover" gorm:"not null"`
	Original    *Song          `json:"original"`
	VideoLink   string         `json:"videoLink"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}
