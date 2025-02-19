package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// AIRequest represents the JSON payload for the AI API
type AIRequest struct {
	PriorityLevel     int     `json:"priority_level"`
	DaysUntilDeadline float32 `json:"days_until_deadline"`
}

// AIResponse represents the response from the AI API
type AIResponse struct {
	PredictedPriority int `json:"predicted_priority"`
}

// PredictPriority calls the Python AI server for task prioritization
func PredictPriority(priorityLevel int, daysUntilDeadline float32) int {
	url := "http://localhost:5000/predict/"

	// Prepare request payload
	requestBody, _ := json.Marshal(AIRequest{
		PriorityLevel:     priorityLevel,
		DaysUntilDeadline: daysUntilDeadline,
	})

	// Make HTTP POST request to AI server
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("❌ Error calling AI server:", err)
		return -1
	}
	defer resp.Body.Close()

	// Decode response
	var aiResponse AIResponse
	json.NewDecoder(resp.Body).Decode(&aiResponse)

	return aiResponse.PredictedPriority
}
