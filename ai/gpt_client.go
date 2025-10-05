package ai

import (
    "context"
    "fmt"
    "github.com/sashabaranov/go-openai"
)

// GetGPTResponse sends the prompt to OpenAI GPT and returns the response.
func GetGPTResponse(client *openai.Client, prompt string) (string, error) {
    ctx := context.Background() // create a context

    // Create the chat completion request
    resp, err := client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
        Model: "gpt-4o-mini",
        Messages: []openai.ChatCompletionMessage{
            {
                Role:    "user",
                Content: prompt,
            },
        },
    })
    if err != nil {
        return "", fmt.Errorf("failed to get GPT response: %w", err)
    }

    if len(resp.Choices) == 0 {
        return "", fmt.Errorf("no response from GPT")
    }

    return resp.Choices[0].Message.Content, nil
}

func NewClient(apiKey string) *openai.Client {
    return openai.NewClient(apiKey)
}
// Example usage
func Example() {
    client := openai.NewClient("OPENAI_API_KEY") // replace with env var ideally
    response, err := GetGPTResponse(client, "Hello, AI!")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("GPT Response:", response)
}
