package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Task struct {
	ID          uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Title       string     `gorm:"not null" json:"title"`
	Description string     `json:"description"`
	Priority    string     `gorm:"type:enum('Low', 'Medium', 'High');default:'Medium'" json:"priority"`
	Status      string     `gorm:"type:enum('Pending', 'In Progress', 'Done');default:'Pending'" json:"status"`
	Deadline    *time.Time `json:"deadline"`
	CreatedBy   uuid.UUID  `gorm:"type:uuid;not null" json:"created_by"`
	CreatedAt   time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	Comments    []Comment  `gorm:"foreignKey:TaskID;constraint:OnDelete:CASCADE;" json:"comments"`
}

// BeforeCreate ensures UUID is generated before inserting a new record
func (t *Task) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = uuid.New()
	return
}
