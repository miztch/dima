package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Handlers struct {
	dbClient *DynamoDBClient
}

func NewHandlers(dbClient *DynamoDBClient) *Handlers {
	return &Handlers{dbClient: dbClient}
}

func (h *Handlers) IndexHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("[info] / handler called")
	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{
		"status":  200,
		"message": "Remember, bullets hurt.",
	}
	json.NewEncoder(w).Encode(response)
}

func (h *Handlers) GetMatchesHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("[info] /matches handler called")
	startDate := time.Now().Format("2006-01-02")

	if inputDate := r.URL.Query().Get("date"); inputDate != "" {
		if IsValidDate(inputDate) {
			startDate = inputDate
		}
	}

	log.Printf("[info] query matches by start date: %s", startDate)
	matches, err := h.dbClient.QueryMatchesByStartDate(context.Background(), startDate)
	if err != nil {
		log.Printf("[error] failed to query matches by start date, %v", err)
		http.Error(w, "failed to query matches", http.StatusInternalServerError)
		return
	}

	log.Printf("[info] found %d matches", len(matches))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(matches)
}
