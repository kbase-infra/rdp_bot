package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// PostMessagePayload represents the payload to send a message to Slack.
type PostMessagePayload struct {
	Channel  string `json:"channel"`
	Text     string `json:"text"`
	ThreadTS string `json:"thread_ts,omitempty"` // Optional thread timestamp
}

func SendMessageToSlack(payload PostMessagePayload, botToken string) error { // Now returns an error

	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshaling JSON: %v", err)
		return err
	}

	req, err := http.NewRequest("POST", "https://slack.com/api/chat.postMessage", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("Error creating HTTP request: %v", err)
		return err
	}

	req.Header.Set("Authorization", "Bearer "+botToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error sending request to Slack: %v", err)
		return err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		return err
	}

	// Log the response status and body for debugging
	log.Printf("Response from Slack: %s", resp.Status)
	log.Printf("Response Body: %s", string(body))

	if resp.StatusCode >= 300 {
		// Handle non-success responses
		log.Printf("Non-success response from Slack: %s", string(body))
		return fmt.Errorf("received non-success status code from Slack: %d", resp.StatusCode)
	}

	log.Printf("Message sent to channel %s", payload.Channel)
	log.Printf("Message: %s", payload.Text)

	return nil // No error occurred
}
