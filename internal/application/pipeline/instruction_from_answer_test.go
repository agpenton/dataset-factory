package pipeline_test

import (
	"context"
	"testing"

	"github.com/agpenton/dataset-factory/internal/application/pipeline"
	"github.com/agpenton/dataset-factory/internal/generator/fake"
)

func TestInstructionFromAnswerPipeline(t *testing.T) {
	g := fake.New("What is Kubernetes?")

	p := pipeline.NewInstructionFromAnswer(g)

	record, err := p.Run(
		context.Background(),
		"Kubernetes is a container orchestration platform.",
	)
	if err != nil {
		t.Fatal(err)
	}

	expected := `{"messages":[{"role":"user","content":"What is Kubernetes?"},{"role":"assistant","content":"Kubernetes is a container orchestration platform."}]}`

	if record != expected {
		t.Fatalf("\nexpected:\n%s\n\ngot:\n%s", expected, record)
	}
}
