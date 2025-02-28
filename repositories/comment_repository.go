package repositories

import (
	"errors"

	"github.com/azka-art/taskwise-backend/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CommentRepository defines all the methods to access comment data
type CommentRepository interface {
	Create(comment *models.Comment) error
	FindByID(id uuid.UUID) (*models.Comment, error)
	FindByTaskID(taskID uuid.UUID) ([]models.Comment, error)
	Update(comment *models.Comment) error
	Delete(id uuid.UUID) error
	DeleteByTaskID(taskID uuid.UUID) error
}

// commentRepository implements the CommentRepository interface
type commentRepository struct {
	db *gorm.DB
}

// NewCommentRepository creates a new comment repository instance
func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &commentRepository{
		db: db,
	}
}

// Create adds a new comment to the database
func (r *commentRepository) Create(comment *models.Comment) error {
	// Validate comment data before saving
	if err := comment.Validate(); err != nil {
		return err
	}

	return r.db.Create(comment).Error
}

// FindByID finds a comment by its ID
func (r *commentRepository) FindByID(id uuid.UUID) (*models.Comment, error) {
	var comment models.Comment

	// Find the comment and preload the User relationship
	err := r.db.Preload("User").Where("id = ?", id).First(&comment).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Return nil, nil when no record is found
		}
		return nil, err
	}

	return &comment, nil
}

// FindByTaskID finds all comments for a specific task
func (r *commentRepository) FindByTaskID(taskID uuid.UUID) ([]models.Comment, error) {
	var comments []models.Comment

	// Find all comments for the task and preload User data
	err := r.db.Preload("User").Where("task_id = ?", taskID).Order("created_at DESC").Find(&comments).Error
	if err != nil {
		return nil, err
	}

	return comments, nil
}

// Update updates an existing comment
func (r *commentRepository) Update(comment *models.Comment) error {
	// Validate comment data before updating
	if err := comment.Validate(); err != nil {
		return err
	}

	// Check if comment exists
	var count int64
	if err := r.db.Model(&models.Comment{}).Where("id = ?", comment.ID).Count(&count).Error; err != nil {
		return err
	}

	if count == 0 {
		return errors.New("comment not found")
	}

	// Update only allowed fields (avoid updating user_id, task_id, etc.)
	return r.db.Model(comment).Updates(map[string]interface{}{
		"content":    comment.Content,
		"updated_at": gorm.Expr("NOW()"),
	}).Error
}

// Delete removes a comment from the database
func (r *commentRepository) Delete(id uuid.UUID) error {
	// Check if comment exists
	var count int64
	if err := r.db.Model(&models.Comment{}).Where("id = ?", id).Count(&count).Error; err != nil {
		return err
	}

	if count == 0 {
		return errors.New("comment not found")
	}

	return r.db.Delete(&models.Comment{}, id).Error
}

// DeleteByTaskID removes all comments for a specific task
func (r *commentRepository) DeleteByTaskID(taskID uuid.UUID) error {
	return r.db.Where("task_id = ?", taskID).Delete(&models.Comment{}).Error
}
