package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Priority represents the priority level of a task
type Priority string

const (
	PriorityLow    Priority = "Low"
	PriorityMedium Priority = "Medium"
	PriorityHigh   Priority = "High"
)

// Status represents the current status of a task
type Status string

const (
	StatusPending    Status = "Pending"
	StatusInProgress Status = "In Progress"
	StatusDone       Status = "Done"
)

// Task represents a task in the system
type Task struct {
	ID          uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Title       string         `gorm:"not null" json:"title"`
	Description string         `json:"description"`
	Priority    Priority       `gorm:"type:enum('Low', 'Medium', 'High');default:'Medium'" json:"priority"`
	Status      Status         `gorm:"type:enum('Pending', 'In Progress', 'Done');default:'Pending'" json:"status"`
	Deadline    *time.Time     `json:"deadline"`
	CreatedBy   uuid.UUID      `gorm:"type:uuid;not null" json:"created_by"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`

	// Relationships
	Comments []Comment `gorm:"foreignKey:TaskID;constraint:OnDelete:CASCADE;" json:"comments,omitempty"`
	User     *User     `gorm:"foreignKey:CreatedBy" json:"user,omitempty"`
}

// BeforeCreate ensures UUID is generated before inserting a new record
func (t *Task) BeforeCreate(tx *gorm.DB) (err error) {
	if t.ID == uuid.Nil {
		t.ID = uuid.New()
	}
	return
}

// Validate checks if the task data is valid
func (t *Task) Validate() error {
	if t.Title == "" {
		return errors.New("title is required")
	}

	if t.CreatedBy == uuid.Nil {
		return errors.New("creator is required")
	}

	// Validate Priority
	if t.Priority != PriorityLow && t.Priority != PriorityMedium && t.Priority != PriorityHigh {
		return errors.New("invalid priority value")
	}

	// Validate Status
	if t.Status != StatusPending && t.Status != StatusInProgress && t.Status != StatusDone {
		return errors.New("invalid status value")
	}

	// Validate Deadline is not in the past
	if t.Deadline != nil {
		now := time.Now()
		if t.Deadline.Before(now) {
			return errors.New("deadline cannot be in the past")
		}
	}

	return nil
}

// SetDefaults sets default values for a new task
func (t *Task) SetDefaults() {
	if t.Priority == "" {
		t.Priority = PriorityMedium
	}

	if t.Status == "" {
		t.Status = StatusPending
	}
}

// TaskRequest represents the data needed to create or update a task
type TaskRequest struct {
	Title       string   `json:"title" binding:"required"`
	Description string   `json:"description"`
	Priority    Priority `json:"priority"`
	Status      Status   `json:"status"`
	Deadline    *string  `json:"deadline"` // Format: "2006-01-02T15:04:05Z"
}

// TaskResponse represents the data returned when a task is requested
type TaskResponse struct {
	ID          uuid.UUID         `json:"id"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Priority    Priority          `json:"priority"`
	Status      Status            `json:"status"`
	Deadline    *time.Time        `json:"deadline"`
	CreatedBy   uuid.UUID         `json:"created_by"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
	Creator     string            `json:"creator,omitempty"`
	Comments    []CommentResponse `json:"comments,omitempty"`
}

// AIRecommendationRequest represents input data for AI task recommendations
type AIRecommendationRequest struct {
	TaskIDs []uuid.UUID `json:"task_ids"`
}
