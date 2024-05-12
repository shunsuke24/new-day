package openai

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"regexp"
	"strings"

	"github.com/new-day/repository"

	"github.com/sashabaranov/go-openai"
	go_openai "github.com/sashabaranov/go-openai"
)

type openAIChat struct {
	client *go_openai.Client
}

func newOpenAIChat(cli *go_openai.Client) repository.OpenAIChat {
	return &openAIChat{
		client: cli,
	}
}

func (r openAIChat) Completion(content, model string, maxTokens int, temperature float32) (*string, error) {
	resp, err := r.client.CreateChatCompletion(
		context.Background(),
		go_openai.ChatCompletionRequest{
			MaxTokens: maxTokens,
			Model:     model,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: content,
				},
			},
			Temperature: temperature,
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return nil, err
	}

	fmt.Println(resp.Choices[0].Message.Content)
	response := resp.Choices[0].Message.Content
	response = strings.Replace(response, "\u3000", " ", -1)
	response = regexp.MustCompile(`\\u3000`).ReplaceAllString(response, " ")
	return &response, nil
}

// written over :
// https://github.com/sashabaranov/go-openai/blob/71a24931dbc5b7029901ff963dc4d0d2509aa7ed/example_test.go#L40C1-L79C2
func (r openAIChat) CompletionStream(conversation []go_openai.ChatCompletionMessage, model string, maxTokens int, temperature float32, c chan string, errs chan error, done chan struct{}) *string {

	stream, err := r.client.CreateChatCompletionStream(
		context.Background(),
		openai.ChatCompletionRequest{
			MaxTokens:   maxTokens,
			Model:       model,
			Messages:    conversation,
			Stream:      true,
			Temperature: temperature,
		},
	)
	if err != nil {
		fmt.Printf("ChatCompletionStream error: %v\n", err)
		errs <- err
		return nil
	}
	defer stream.Close()
	defer close(c)
	defer close(errs)
	defer close(done)

	for {
		var response go_openai.ChatCompletionStreamResponse
		response, err = stream.Recv()
		if errors.Is(err, io.EOF) {
			done <- struct{}{} //終了通知
			return nil
		}

		if err != nil {
			log.Printf("Stream error: %v\n", err)
			errs <- err
			return nil
		}

		streamingContent := response.Choices[0].Delta.Content
		streamingContent = strings.Replace(streamingContent, "\u3000", " ", -1)
		streamingContent = regexp.MustCompile(`\\u3000`).ReplaceAllString(streamingContent, " ")
		c <- streamingContent
	}
}
