package openai

import (
	"github.com/new-day/repository"
	go_openai "github.com/sashabaranov/go-openai"
)

func NewOpenAIRepository(client *go_openai.Client) *repository.OpenAI {
	return &repository.OpenAI{
		OpenAIChat: newOpenAIChat(client),
	}
}
