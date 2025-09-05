package logic

import (
	"context"
	"fmt"
	"os"

	"github.com/sashabaranov/go-openai"
)

func GenerateAISummary(income, expenses, surplus, score float64) (string, error) {
	key := os.Getenv("OPENAI_API_KEY")
	if key == "" {
		return "No API key provided. Skipping AI summary.", nil
	}

	client := openai.NewClient(key)
	prompt := fmt.Sprintf("Income: £%.2f, Expenses: £%.2f, Surplus: £%.2f, Health Score: %.2f%%. Write a short financial summary for the user.", income, expenses, surplus, score)

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{Role: "user", Content: prompt},
			},
		},
	)
	if err != nil {
		return "", err
	}
	return resp.Choices[0].Message.Content, nil
}
