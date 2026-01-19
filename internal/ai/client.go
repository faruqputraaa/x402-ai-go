package ai

import (
	"context"
	"errors"

	"google.golang.org/genai"
)

type Client struct {
	client *genai.Client
	model  string
}

func New(apiKey string) (*Client, error) {
	ctx := context.Background()

	c, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey: apiKey,
	})
	if err != nil {
		return nil, err
	}

	model := "gemini-3-flash-preview" // atau gemini-1.5-flash

	return &Client{
		client: c,
		model:  model,
	}, nil
}

func (c *Client) Generate(prompt string) (string, error) {
	ctx := context.Background()

	result, err := c.client.Models.GenerateContent(
		ctx,
		c.model,
		genai.Text(prompt),
		nil,
	)
	if err != nil {
		return "", err
	}

	text := result.Text()
	if text == "" {
		return "", errors.New("gemini: empty response")
	}

	return text, nil
}
