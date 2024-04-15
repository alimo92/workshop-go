package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	slackHandler "workshop/internal"

	"github.com/slack-go/slack"
)

func ActionHandler(w http.ResponseWriter, r *http.Request) {
	var payload slack.InteractionCallback
	err := json.Unmarshal([]byte(r.FormValue("payload")), &payload)
	if err != nil {
		fmt.Printf("Could not parse action response JSON: %v\n", err)
	}

	slackHandler.SendMessage(fmt.Sprintf("Message button pressed by user %s with value %s", payload.User.Name, payload.ActionCallback.AttachmentActions[0].Value))
}

func ActionHandlerV1(w http.ResponseWriter, r *http.Request) {


	fmt.Print(io.ReadAll(r.Body))
}
