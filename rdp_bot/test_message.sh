#!/usr/bin/env bash
# This script sends a test message to a locally running bot
export TOKEN=your_slack_bot_token_here

curl -X POST http://localhost:8081/slack/events \
                           -H 'Content-Type: application/json' \
                           -d '{
                           "token": "$TOKEN",
                           "team_id": "T0001",
                           "api_app_id": "A0001",
                           "event": {
                               "type": "app_mention",
                               "user": "U061F7AUR",
                               "text": "<@U0LAN0Z89> is it going to rain tomorrow?",
                               "ts": "1515449522.000016",
                               "channel": "C0LAN2Q65",
                               "event_ts": "1515449522000016"
                           },
                           "type": "event_callback",
                           "authed_teams": ["T0001"],
                           "event_id": "Ev0LAN670R",
                           "event_time": 1515449522000016
                       }'