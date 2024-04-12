# How to run the script to get the list of RDP users
1. Open the terminal
2. Run the following command:
```bash
export SLACK_TOKEN_XOXB='your_slack_bot_token_here'
export USER_EMAIL_CSV='abc@gmail.com,abc123@gmail.com'
go run main.go
```
3. The list of RDP users will be displayed in the terminal
```
Emails CSV: abc@gmail.com,abc123@gmail.com
Fetched RDP User IDs: [UQM22X3N,UQM22X32  ]
```