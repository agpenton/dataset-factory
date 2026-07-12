package prompt_test

import (
	"testing"

	"github.com/agpenton/dataset-factory/internal/application/prompt"
)

func TestBuildInstructionFromAnswer(t *testing.T) {
	answer := "Kubernetes is a container orchestration platform."

	got := prompt.BuildInstructionFromAnswer(answer)

	expected := `Generate the user instruction that would produce the following answer:

Kubernetes is a container orchestration platform.`

	if got != expected {
		t.Fatalf("unexpected prompt:\n%s", got)
	}
}
