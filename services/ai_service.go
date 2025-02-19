package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

// AIRequest merepresentasikan request ke AI server
type AIRequest struct {
	PriorityLevel     int     `json:"priority_level"`
	DaysUntilDeadline float32 `json:"days_until_deadline"`
}

// PredictTaskPriority mengirimkan request ke server AI dan menerima respons
func GetAIPrediction(priorityLevel int, daysUntilDeadline float32) (int, error) {
	aiURL := "http://localhost:5000/predict/"
	payload := AIRequest{PriorityLevel: priorityLevel, DaysUntilDeadline: daysUntilDeadline}
	jsonPayload, _ := json.Marshal(payload)

	// Kirim request POST ke AI server
	resp, err := http.Post(aiURL, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return 0, errors.New("AI server unavailable")
	}
	defer resp.Body.Close()

	// Parsing respons AI
	var response struct {
		PredictedPriority int `json:"predicted_priority"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return 0, err
	}

	return response.PredictedPriority, nil
}
