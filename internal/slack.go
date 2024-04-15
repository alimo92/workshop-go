package slack

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func SendMessage(value string) {
	var token, channel string
	var ok bool

	token, ok = os.LookupEnv("SLACK_TOKEN")
	if !ok {
		fmt.Println("Missing SLACK_TOKEN in environment")
		os.Exit(1)
	}

	channel, ok = os.LookupEnv("SLACK_CHANNEL")
	if !ok {
		fmt.Println("Missing SLACK_CHANNEL in environment")
		os.Exit(1)
	}

	api := slack.New(token)

	api.PostMessage(channel, slack.MsgOptionText(value, false))
}

func SendButtonMessage() {
	var token, channel string
	var ok bool

	token, ok = os.LookupEnv("SLACK_TOKEN")
	if !ok {
		fmt.Println("Missing SLACK_TOKEN in environment")
		os.Exit(1)
	}

	channel, ok = os.LookupEnv("SLACK_CHANNEL")
	if !ok {
		fmt.Println("Missing SLACK_CHANNEL in environment")
		os.Exit(1)
	}

	api := slack.New(token)
	attachment := slack.Attachment{
		Pretext:    "pretext",
		Fallback:   "We don't currently support your client",
		CallbackID: "accept_or_reject",
		Color:      "#3AA3E3",
		Actions: []slack.AttachmentAction{
			slack.AttachmentAction{
				Name:  "accept",
				Text:  "Accept",
				Type:  "button",
				Value: "accept",
			},
			slack.AttachmentAction{
				Name:  "reject",
				Text:  "Reject",
				Type:  "button",
				Value: "reject",
				Style: "danger",
			},
		},
	}

	message := slack.MsgOptionAttachments(attachment)
	channelID, timestamp, err := api.PostMessage(channel, slack.MsgOptionText("", false), message)
	if err != nil {
		fmt.Printf("Could not send message: %v\n", err)
	}
	fmt.Printf("Message with buttons sucessfully sent to channel %s at %s\n", channelID, timestamp)
}
