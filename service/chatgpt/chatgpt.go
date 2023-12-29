package chatgpt

import (
	"context"
	"fmt"

	openai "github.com/sashabaranov/go-openai"
)

func SendMessage(message string) (string, error) {
	client := openai.NewClient("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")

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: message,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return "", err
	}
	return resp.Choices[0].Message.Content, nil
}
