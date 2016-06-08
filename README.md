# DeployBot
This is a bot that gets incoming webhooks (from slack first) and can deploy a Go app to a Google App Engine Flexible VM

## Configuration
You'll have to provide DeployBot with the correct token in order to validate the incoming hooks from Slack

## Building
Make sure you're using Go 1.6 or Go 1.5 with the GO15VENDOREXPERIMENT variable set to 1

```
$ go build *.go -o deploybot

```

## Running

You'll need to set the following ENV for DeployBot to be able to authenticate incoming hooks
```
SLACK_TOKEN     you get this from Slack when you create the Outgoing Webhooks
SLACK_TEAM      this is the name of the team, not the ID
SLACK_CHANNEL   this is the human name for the channel the hook will be allowed from, not the ID
```

Export those and then just run the bot

```
$ ./deploybot
```
