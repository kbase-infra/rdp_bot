package util

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

// SlackUserLookupResponse represents the structure of the response from Slack's users.lookupByEmail API.
type SlackUserLookupResponse struct {
	User struct {
		ID string `json:"id"`
	} `json:"user"`
}

// GetRdpUsers returns a slice of user IDs. It looks up Slack user IDs for the given email addresses if provided,
// otherwise, it uses the provided list of user IDs.
func GetRdpUsers(emailsCSV, usersCSV, xoxbToken string) ([]string, error) {
	// If the list of user IDs is provided, use it directly.
	if usersCSV != "" {
		userIDs := strings.Split(usersCSV, ",")
		return userIDs, nil
	}

	// If the list of emails is provided, perform the lookup.
	if emailsCSV != "" {
		var userIDs []string
		emails := strings.Split(emailsCSV, ",")
		for _, email := range emails {
			userID, err := lookupUserIDByEmail(email, xoxbToken)
			if err != nil {
				log.Fatalf("Failed to lookup user ID for email %s: %v", email, err) // This will log the error and terminate the program.
			}
			userIDs = append(userIDs, userID)
		}
		return userIDs, nil
	}

	return nil, fmt.Errorf("either USER_EMAIL_CSV or USER_ID_CSV must be provided")
}

// lookupUserIDByEmail makes a call to Slack's users.lookupByEmail API to find the user ID associated with the given email address.
func lookupUserIDByEmail(email, xoxbToken string) (string, error) {
	if xoxbToken == "" {
		log.Fatalf("xoxb_token is required to lookup user ID by email")
	}

	req, err := http.NewRequest("GET", "https://slack.com/api/users.lookupByEmail", nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	q := req.URL.Query()
	q.Add("email", email)
	req.URL.RawQuery = q.Encode()

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", xoxbToken))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("request failed: %w", err)
	}
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			log.Printf("Failed to close response body: %v", closeErr)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("slack API responded with status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body) // Replacing ioutil.ReadAll with io.ReadAll
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	var lookupResponse SlackUserLookupResponse
	if err := json.Unmarshal(body, &lookupResponse); err != nil {
		return "", fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return lookupResponse.User.ID, nil
}
