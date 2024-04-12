# Random Devops Person Bot (rdp_bot)




# Push to your docker registry
```
docker build --platform=linux/amd64 . -t <org>/rdp_bot:latest
docker push <org>/devops_poc_bot:latest
```
## Host it at your URL

* configure your manifest.yaml with that hosted url

## Deployment

* SLACK_TOKEN_XOXB = xoxb-BotTokenGoesHere
* USER_EMAIL_CSV = csv of usernames in the rotation, e.g email1@anl.gov,email2@anl.gov 

## NGINX
```
        location ~ /services/(rdpbot)/(.*){
                set $servicehost rdpbot;
                set $serviceurl $2;
                proxy_pass http://$servicehost:8081/$serviceurl;
                proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
                proxy_set_header X-Real-IP $remote_addr;
                proxy_set_header Host $http_host;
                proxy_set_header X-Forwarded-Proto $scheme;
                more_set_headers "Access-Control-Allow-Headers: authorization, Content-Type";
        }
```

## Running Locally
```
go build -v -o rdp_bot
export SLACK_TOKEN_XOXB=your_slack_bot_token_here
export USER_EMAIL_CSV=email@email.com,email2@email.com
./rdp_bot
```
