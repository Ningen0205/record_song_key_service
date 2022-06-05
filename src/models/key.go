package model

import (
	"time"

	"gorm.io/gorm"
)

type Key struct {
	ID        int            `json:"id" gorm:"primarykey"`
	Value     int            `json:"value" gorm:"unique;not null"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}
