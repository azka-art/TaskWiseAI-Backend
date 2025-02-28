package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/azka-art/taskwise-backend/models"
	"github.com/azka-art/taskwise-backend/repositories"
	"github.com/google/uuid"
)

// AIRequest merepresentasikan request ke AI server
type AIRequest struct {
	PriorityLevel     int     `json:"priority_level"`
	DaysUntilDeadline float32 `json:"days_until_deadline"`
}

// AIResponse merepresentasikan respons dari AI server
type AIResponse struct {
	PredictedPriority int `json:"predicted_priority"`
}

// GetAIPrediction mengirimkan request ke server AI dan menerima respons
func GetAIPrediction(priorityLevel int, daysUntilDeadline float32) (int, error) {
	// Get AI server URL from environment or use default
	aiURL := os.Getenv("AI_SERVER_URL")
	if aiURL == "" {
		aiURL = "http://localhost:5000/predict/"
	}

	// Create request payload
	payload := AIRequest{
		PriorityLevel:     priorityLevel,
		DaysUntilDeadline: daysUntilDeadline,
	}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return 0, fmt.Errorf("error creating request payload: %w", err)
	}

	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	// Kirim request POST ke AI server
	resp, err := client.Post(aiURL, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return 0, errors.New("AI server unavailable")
	}
	defer resp.Body.Close()

	// Check status code
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("AI server returned status: %d", resp.StatusCode)
	}

	// Parsing respons AI
	var response AIResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return 0, fmt.Errorf("error parsing AI response: %w", err)
	}

	return response.PredictedPriority, nil
}

// ConvertPriorityToLevel converts Priority type to numeric level
func ConvertPriorityToLevel(priority models.Priority) int {
	switch priority {
	case models.PriorityLow:
		return 1
	case models.PriorityHigh:
		return 3
	default:
		return 2 // Medium is default
	}
}

// ConvertLevelToPriority converts numeric level to Priority type
func ConvertLevelToPriority(level int) models.Priority {
	switch level {
	case 1:
		return models.PriorityLow
	case 3:
		return models.PriorityHigh
	default:
		return models.PriorityMedium
	}
}

// GetTaskPriorities generates AI recommended priorities for a list of tasks
func GetTaskPriorities(taskIDs []uuid.UUID) (map[uuid.UUID]models.Priority, error) {
	result := make(map[uuid.UUID]models.Priority)

	// Return empty result if no task IDs provided
	if len(taskIDs) == 0 {
		return result, nil
	}

	// Get all tasks
	for _, taskID := range taskIDs {
		task, err := repositories.GetTaskByID(taskID)
		if err != nil || task == nil {
			// Skip tasks that can't be found
			continue
		}

		// Calculate days until deadline
		var daysUntilDeadline float32 = 30 // Default if no deadline
		if task.Deadline != nil {
			daysUntilDeadline = float32(time.Until(*task.Deadline).Hours() / 24)
		}

		// Get priority level from task
		priorityLevel := ConvertPriorityToLevel(task.Priority)

		// Get AI prediction
		predictedLevel, err := GetAIPrediction(priorityLevel, daysUntilDeadline)
		if err != nil {
			// Log error but continue with other tasks
			fmt.Printf("Error getting AI prediction for task %s: %v\n", task.ID, err)
			continue
		}

		// Convert predicted level back to Priority type
		predictedPriority := ConvertLevelToPriority(predictedLevel)

		// Store result
		result[task.ID] = predictedPriority
	}

	return result, nil
}

// PrioritizeTasks sorts a list of tasks based on AI recommendations
func PrioritizeTasks(tasks []models.Task) ([]models.Task, error) {
	if len(tasks) == 0 {
		return tasks, nil
	}

	// Create a copy of tasks to avoid modifying the original
	tasksCopy := make([]models.Task, len(tasks))
	copy(tasksCopy, tasks)

	// Collect task IDs
	taskIDs := make([]uuid.UUID, len(tasksCopy))
	for i, task := range tasksCopy {
		taskIDs[i] = task.ID
	}

	// Get AI recommendations
	recommendations, err := GetTaskPriorities(taskIDs)
	if err != nil {
		return nil, err
	}

	// Apply recommendations
	for i, task := range tasksCopy {
		if priority, exists := recommendations[task.ID]; exists {
			tasksCopy[i].Priority = priority
		}
	}

	// Sort tasks into priority buckets
	var highPriority, mediumPriority, lowPriority []models.Task

	for _, task := range tasksCopy {
		switch task.Priority {
		case models.PriorityHigh:
			highPriority = append(highPriority, task)
		case models.PriorityMedium:
			mediumPriority = append(mediumPriority, task)
		case models.PriorityLow:
			lowPriority = append(lowPriority, task)
		}
	}

	// Combine buckets in priority order
	prioritizedTasks := make([]models.Task, 0, len(tasksCopy))
	prioritizedTasks = append(prioritizedTasks, highPriority...)
	prioritizedTasks = append(prioritizedTasks, mediumPriority...)
	prioritizedTasks = append(prioritizedTasks, lowPriority...)

	return prioritizedTasks, nil
}
