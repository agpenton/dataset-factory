package input_test

import (
	"testing"

	"github.com/agpenton/dataset-factory/internal/application/input"
)

func TestReadAnswers(t *testing.T) {
	records, err := input.ReadAnswers("testdata/answers.jsonl")
	if err != nil {
		t.Fatal(err)
	}

	if len(records) != 2 {
		t.Fatalf("expected 2 records, got %d", len(records))
	}

	if records[0].Answer != "Kubernetes is a container orchestration platform." {
		t.Fatal("unexpected first answer")
	}

	if records[1].Answer != "Docker packages applications into containers." {
		t.Fatal("unexpected second answer")
	}
}
