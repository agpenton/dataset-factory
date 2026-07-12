package ollama_test

import (
	"context"
	"os"
	"testing"

	"github.com/agpenton/dataset-factory/internal/generator/ollama"
)

func TestGenerateIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	if os.Getenv("OLLAMA_TEST") == "" {
		t.Skip("OLLAMA_TEST not set")
	}

	client := ollama.New(
		ollama.DefaultEndpoint,
		ollama.DefaultModel,
	)

	answer, err := client.Generate(
		context.Background(),
		"Say hello.",
	)
	if err != nil {
		t.Fatal(err)
	}

	if answer == "" {
		t.Fatal("expected non-empty response")
	}
}
