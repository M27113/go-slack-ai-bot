package handler

import (
	"fmt"
	"log"
	"os"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
	"github.com/slack-go/slack/slackevents"

	"go-slack-ai-bot/ai"
)

// StartSlackBot runs the Slack bot in Socket Mode
func StartSlackBot() {
	appToken := os.Getenv("SLACK_APP_TOKEN")
	botToken := os.Getenv("SLACK_BOT_TOKEN")

	if appToken == "" || botToken == "" {
		log.Fatal("SLACK_APP_TOKEN or SLACK_BOT_TOKEN is not set")
	}

	client := slack.New(botToken,
		slack.OptionDebug(true),
		slack.OptionAppLevelToken(appToken),
	)

	socketClient := socketmode.New(
		client,
		socketmode.OptionDebug(true),
	)

	fmt.Println("Slack AI Bot started...")

	go func() {
		for evt := range socketClient.Events {
			switch evt.Type {

			case socketmode.EventTypeEventsAPI:
				eventsAPIEvent, ok := evt.Data.(slackevents.EventsAPIEvent)
				if !ok {
					log.Println("Ignored: not an EventsAPIEvent")
					continue
				}

				socketClient.Ack(*evt.Request)

				if eventsAPIEvent.Type == slackevents.CallbackEvent {
					switch ev := eventsAPIEvent.InnerEvent.Data.(type) {

					// Handle mentions in channels
					case *slackevents.AppMentionEvent:
						go handleMention(ev, client)

					// Handle direct messages
					case *slackevents.MessageEvent:
						// Ignore messages from the bot itself
						authResp, err := client.AuthTest()
						if err != nil {
							log.Printf("Failed to get bot user ID: %v", err)
							break
						}
						botUserID := authResp.UserID
						if ev.User == "" || ev.User == botUserID {
							break
						}
						go handleDirectMessage(ev, client)
					}
				}

			case socketmode.EventTypeInteractive:
				// future: handle buttons/slash commands
			}
		}
	}()

	err := socketClient.Run()
	if err != nil {
		log.Fatalf("Failed to run socketmode client: %v", err)
	}
}

// handleMention responds to @mentions in channels
func handleMention(event *slackevents.AppMentionEvent, client *slack.Client) {
	prompt := event.Text
	channel := event.Channel
	user := event.User

	openAIClient := ai.NewClient(os.Getenv("OPENAI_API_KEY"))

	respText, err := ai.GetGPTResponse(openAIClient, prompt)
	if err != nil {
		respText = fmt.Sprintf("Error getting response: %v", err)
	}

	_, _, err = client.PostMessage(channel,
		slack.MsgOptionText(fmt.Sprintf("<@%s> %s", user, respText), false))
	if err != nil {
		log.Printf("Failed to post message: %v", err)
	}
}

// handleDirectMessage responds to direct messages
func handleDirectMessage(event *slackevents.MessageEvent, client *slack.Client) {
	prompt := event.Text
	channel := event.Channel

	openAIClient := ai.NewClient(os.Getenv("OPENAI_API_KEY"))

	respText, err := ai.GetGPTResponse(openAIClient, prompt)
	if err != nil {
		respText = fmt.Sprintf("Error: %v", err)
	}

	_, _, err = client.PostMessage(channel, slack.MsgOptionText(respText, false))
	if err != nil {
		log.Printf("Failed to post DM: %v", err)
	}
}
