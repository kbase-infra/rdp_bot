package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"rdp_bot/util"
	"time"
)

// SlackEvent represents the structure of the event data from Slack
type SlackEvent struct {
	Token     string                 `json:"token"`
	TeamID    string                 `json:"team_id"`
	APIAppID  string                 `json:"api_app_id"`
	Event     map[string]interface{} `json:"event"`
	Type      string                 `json:"type"`
	EventID   string                 `json:"event_id"`
	EventTime int64                  `json:"event_time"`
	Challenge string                 `json:"challenge,omitempty"` // Add this line
}

func HandleSlackEvents(w http.ResponseWriter, r *http.Request, users []string, xoxbToken string) { // Add users parameter

	// Check if the request method is GET and return a JSON message
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "application/json")
		response := map[string]string{
			"message":           "This is a GET request. Slack makes POST requests.",
			"xoxb_token_status": "not_yet_implemented",
		}
		if err := json.NewEncoder(w).Encode(response); err != nil {
			log.Printf("Error writing response: %v", err)
			http.Error(w, "Error generating response", http.StatusInternalServerError)
		}
		return
	}

	var event SlackEvent
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		log.Printf("Error decoding request body: %v", err)
		log.Printf("Request body: %v", r.Body)
		return
	}

	// Respond to URL verification challenge from Slack
	if event.Type == "url_verification" {
		w.Header().Set("Content-Type", "text/plain")
		_, err := w.Write([]byte(event.Challenge))
		if err != nil {
			log.Printf("Error writing challenge response: %v", err)
		} else {
			log.Printf("Challenge response sent: %s", event.Challenge)
		}
		return
	}

	// Handle event callback
	if event.Type == "event_callback" {
		if eventType, ok := event.Event["type"].(string); ok && eventType == "app_mention" {
			handleAppMention(event, users, xoxbToken)
		}
	}
}

func handleAppMention(event SlackEvent, users []string, xoxbToken string) {
	channel := event.Event["channel"].(string)
	threadTS := event.Event["ts"].(string) // Use the original message timestamp as thread_ts to reply in thread

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	selectedUserID := users[r.Intn(len(users))] // Pick a random user

	// Construct the message payload
	payload := util.PostMessagePayload{
		Channel:  channel,
		Text:     fmt.Sprintf("<@%s> is the next user.", selectedUserID),
		ThreadTS: threadTS,
	}

	util.SendMessageToSlack(payload, xoxbToken)
}
