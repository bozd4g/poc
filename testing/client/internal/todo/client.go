package todo

import (
	"context"
	"fmt"

	gohttpclient "github.com/bozd4g/go-http-client"
)

type Client struct {
	httpClient *gohttpclient.Client
}

func NewClient(url string) *Client {
	return &Client{gohttpclient.New(url)}
}

func (c *Client) GetTodos(ctx context.Context) ([]TodoDto, error) {
	res, err := c.httpClient.Get(ctx, "/api/todos")
	if err != nil {
		return nil, fmt.Errorf("failed to get todos: %w", err)
	}

	fmt.Println(string(res.Body()))
	var dtos []TodoDto
	if err := res.Unmarshal(&dtos); err != nil {
		return nil, fmt.Errorf("failed to unmarshal todos: %w", err)
	}

	return dtos, nil
}
