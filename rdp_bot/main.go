package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"rdp_bot/handlers"
	"rdp_bot/util" // Import the util package to use the get_rdp_users function
	"strconv"
)

func checkRequiredEnvVars() {
	// Check for Slack tokens in SLACK_TOKEN_XOXB
	if os.Getenv("SLACK_TOKEN_XOXB") == "" {
		log.Fatalf("SLACK_TOKEN_XOXB environment variable must be set")
	}
	// Check that exactly one of USER_EMAIL_CSV or USER_ID_CSV is set
	userEmailCSV := os.Getenv("USER_EMAIL_CSV")
	userIDCSV := os.Getenv("USER_ID_CSV")

	if userEmailCSV == "" && userIDCSV == "" {
		log.Fatalf("One of USER_EMAIL_CSV or USER_ID_CSV environment variable must be set")
	}

	if userEmailCSV != "" && userIDCSV != "" {
		log.Fatalf("Both USER_EMAIL_CSV and USER_ID_CSV environment variables are set, only one is allowed")
	}
}

func main() {
	checkRequiredEnvVars()

	users, err := util.GetRdpUsers(os.Getenv("USER_EMAIL_CSV"), os.Getenv("USER_ID_CSV"), os.Getenv("SLACK_TOKEN_XOXB"))
	if err != nil {
		// Print out the error and exit
		log.Fatalf("Failed to fetch RDP users: %v", err)
	}

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8081" // Default port if not specified
	}

	// Validate if port is a number
	if _, err := strconv.Atoi(port); err != nil {
		log.Fatalf("Invalid SERVER_PORT: %v", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.HandleSlackEvents(w, r, users, os.Getenv("SLACK_TOKEN_XOXB"))
	})

	fmt.Printf("Server started on :%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
