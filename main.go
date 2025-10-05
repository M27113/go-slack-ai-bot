package main

import (
	"fmt"
	"os"

	"go-slack-ai-bot/handler"
)

func main() {
	// Ensure required environment variables are set
	requiredEnv := []string{"SLACK_BOT_TOKEN", "SLACK_APP_TOKEN", "OPENAI_API_KEY"}
	for _, env := range requiredEnv {
		if os.Getenv(env) == "" {
			fmt.Printf("Error: %s environment variable is not set\n", env)
			return
		}
	}

	// Start the Slack bot
	handler.StartSlackBot()
}
