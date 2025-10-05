package config

import "os"

type Config struct {
    SlackBotToken  string
    SlackAppToken  string
    OpenAIKey      string
}

func LoadConfig() *Config {
    return &Config{
        SlackBotToken: os.Getenv("SLACK_BOT_TOKEN"),
        SlackAppToken: os.Getenv("SLACK_APP_TOKEN"),
        OpenAIKey:     os.Getenv("OPENAI_API_KEY"),
    }
}
