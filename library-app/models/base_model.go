package models

import (
	"github.com/google/uuid"
	"time"
)

type BaseModel struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime;"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime;"`
}
