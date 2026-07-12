package ollama

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	endpoint string
	model    string
	client   *http.Client
}

const (
	DefaultEndpoint = "http://localhost:11434"
	DefaultModel    = "qwen3:8b-q4_K_M"
)

func New(endpoint, model string) *Client {
	return &Client{
		endpoint: endpoint,
		model:    model,
		client:   http.DefaultClient,
	}
}

func (c *Client) Generate(
	ctx context.Context,
	prompt string,
) (string, error) {

	reqBody := request{
		Model:  c.model,
		Prompt: prompt,
		Stream: false,
	}

	body, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		c.endpoint+"/api/generate",
		bytes.NewReader(body),
	)
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)

		return "", fmt.Errorf(
			"unexpected status %s: %s",
			resp.Status,
			string(body),
		)
	}

	var result response

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	return result.Response, nil
}
