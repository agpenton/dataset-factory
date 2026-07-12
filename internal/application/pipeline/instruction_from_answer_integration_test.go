package pipeline_test

import (
	"context"
	"testing"

	"github.com/agpenton/dataset-factory/internal/application/input"
	"github.com/agpenton/dataset-factory/internal/application/pipeline"
	"github.com/agpenton/dataset-factory/internal/generator/fake"
)

func TestGenerateDatasetFromAnswers(t *testing.T) {
	records, err := input.ReadAnswers("../input/testdata/answers.jsonl")
	if err != nil {
		t.Fatal(err)
	}

	g := fake.New("What is Kubernetes?")

	p := pipeline.NewInstructionFromAnswer(g)

	output, err := p.RunAll(context.Background(), records)
	if err != nil {
		t.Fatal(err)
	}

	if len(output) != 2 {
		t.Fatalf("expected 2 records, got %d", len(output))
	}
}
