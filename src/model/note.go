package model

import (
	"time"

	"gorm.io/gorm"
)

type Note struct {
	ID          int            `json:"id" gorm:"primarykey"`
	Description int            `json:"description" gorm:"not null"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}
