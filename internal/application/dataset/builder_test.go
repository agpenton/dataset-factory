package dataset_test

import (
	"encoding/json"
	"testing"

	"github.com/agpenton/dataset-factory/internal/application/dataset"
)

func TestBuildChatMLRecord(t *testing.T) {
	record := dataset.BuildRecord(
		"What is Kubernetes?",
		"Kubernetes is a container orchestration platform.",
	)

	var decoded map[string]any

	if err := json.Unmarshal([]byte(record), &decoded); err != nil {
		t.Fatal(err)
	}

	messages := decoded["messages"].([]any)

	if len(messages) != 2 {
		t.Fatalf("expected 2 messages, got %d", len(messages))
	}
}
