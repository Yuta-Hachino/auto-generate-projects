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
