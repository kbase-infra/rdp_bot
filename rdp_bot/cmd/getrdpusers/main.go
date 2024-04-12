package main

import (
	"fmt"
	"log"
	"os"
	"rdp_bot/util" // Use the correct import path based on your module's name
)

func main() {
	emailsCSV := os.Getenv("USER_EMAIL_CSV")
	usersCSV := os.Getenv("USER_ID_CSV")
	xoxb_token := os.Getenv("SLACK_XOXB_TOKEN")

	userIDs, err := util.GetRdpUsers(emailsCSV, usersCSV, xoxb_token)
	if err != nil {
		log.Fatalf("Error fetching RDP user IDs: %v", err)
	}
	// print out emails and fetched user IDs
	fmt.Printf("Emails CSV: %s\n", emailsCSV)
	fmt.Println("Fetched RDP User IDs:", userIDs)
}
