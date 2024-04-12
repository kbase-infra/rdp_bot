#!/usr/bin/env bash
# This script builds the bot and runs it
go build -v -o rdp_bot
export SLACK_TOKEN_XOXB=your_slack_bot_token_here
export USER_EMAIL_CSV=email@email.com,email2@email.com
./rdp_bot
