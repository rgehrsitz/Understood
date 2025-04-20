package summarizer

import (
	"context"
	"fmt"
	"os"
	"io/ioutil"

	openai "github.com/sashabaranov/go-openai"
)

type Summarizer struct {
	Client *openai.Client
	Model  string
}

// SummarizeFile reads file at path, sends it to LLM, returns summary.
func (s *Summarizer) SummarizeFile(ctx context.Context, path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("reading %q: %w", path, err)
	}
	prompt := fmt.Sprintf(`You are a code documentation assistant.\nGiven the contents of file %s, produce:\n1. A one-sentence summary of its purpose.\n2. A bullet list of its exported functions/types and what they do.\n3. Any notable implementation details or patterns.\n\n\u0060\u0060\u0060go\n%s\n\u0060\u0060\u0060\n`, path, string(data))
	resp, err := s.Client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model: s.Model,
		Messages: []openai.ChatCompletionMessage{
			{Role: "system", Content: "You generate concise, accurate code documentation."},
			{Role: "user", Content: prompt},
		},
		Temperature: 0.2,
	})
	if err != nil {
		return "", fmt.Errorf("LLM summarization error for %q: %w", path, err)
	}
	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("no summary returned for %q", path)
	}
	return resp.Choices[0].Message.Content, nil
}

// NewSummarizer creates a new Summarizer with the given model.
func NewSummarizer(model string) *Summarizer {
	apiKey := os.Getenv("OPENAI_API_KEY")
	return &Summarizer{
		Client: openai.NewClient(apiKey),
		Model:  model,
	}
}
