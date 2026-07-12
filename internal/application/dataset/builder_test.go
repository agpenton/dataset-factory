package dataset_test

import (
	"encoding/json"
	"fmt"
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

func TestBuildInstructionPrompt(t *testing.T) {
	answer := "Kubernetes is a container orchestration platform."

	prompt := prompt.BuildInstructionFromAnswer(answer)

	expected := `Generate the user instruction that would produce the following answer:

Kubernetes is a container orchestration platform.`

	if prompt != expected {
		t.Fatalf("unexpected prompt:\n%s", prompt)
	}
}

func BuildInstructionFromAnswer(answer string) string {
	return fmt.Sprintf(
		"Generate the user instruction that would produce the following answer:\n\n%s",
		answer,
	)
}
