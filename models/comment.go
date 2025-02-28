package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Comment represents a comment on a task
type Comment struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	TaskID    uuid.UUID      `gorm:"type:uuid;not null" json:"task_id"`
	UserID    uuid.UUID      `gorm:"type:uuid;not null" json:"user_id"`
	Content   string         `gorm:"not null" json:"content"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Add these if you want to include related data in JSON responses
	User *User `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Task *Task `gorm:"foreignKey:TaskID" json:"task,omitempty"`
}

// BeforeCreate ensures UUID is generated before inserting a new record
func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	if c.ID == uuid.Nil {
		c.ID = uuid.New()
	}
	return
}

// Validate checks if the comment data is valid
func (c *Comment) Validate() error {
	if c.TaskID == uuid.Nil {
		return errors.New("task ID is required")
	}

	if c.UserID == uuid.Nil {
		return errors.New("user ID is required")
	}

	if c.Content == "" {
		return errors.New("comment content is required")
	}

	return nil
}

// CommentRequest represents the data needed to create or update a comment
type CommentRequest struct {
	Content string `json:"content" binding:"required"`
}

// CommentResponse represents the data returned when a comment is requested
type CommentResponse struct {
	ID        uuid.UUID `json:"id"`
	TaskID    uuid.UUID `json:"task_id"`
	UserID    uuid.UUID `json:"user_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Username  string    `json:"username,omitempty"`
}
