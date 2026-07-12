package ollama_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/agpenton/dataset-factory/internal/generator/ollama"
)

func TestGenerate(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		w.Write([]byte(`{
			"response":"What is Kubernetes?",
			"done":true
		}`))
	}))
	defer server.Close()

	client := ollama.New(server.URL, "qwen3:8b")

	answer, err := client.Generate(
		context.Background(),
		"ignored",
	)
	if err != nil {
		t.Fatal(err)
	}

	if answer != "What is Kubernetes?" {
		t.Fatal(answer)
	}
}
