package repository

import (
	go_openai "github.com/sashabaranov/go-openai"
)

type OpenAI struct {
	OpenAIChat OpenAIChat
}

type OpenAIChat interface {
	Completion(input, model string, maxTokens int, temperature float32) (*string, error) // deprecated
	CompletionStream(conversation []go_openai.ChatCompletionMessage, model string, maxTokens int, temperature float32, c chan string, errs chan error, done chan struct{}) *string
}
