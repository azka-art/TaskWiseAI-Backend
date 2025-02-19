package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Comment struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	TaskID    uuid.UUID `gorm:"type:uuid;not null" json:"task_id"`
	UserID    uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	Content   string    `gorm:"not null" json:"content"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}

// BeforeCreate ensures UUID is generated before inserting a new record
func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New()
	return
}
