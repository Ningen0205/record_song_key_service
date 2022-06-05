package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID          uuid.UUID      `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	LowestNote  Note           `json:"lowestNote"`
	HighestNote Note           `json:"highestNote"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}
