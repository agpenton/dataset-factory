package run_test

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/agpenton/dataset-factory/internal/application/run"
	"github.com/agpenton/dataset-factory/internal/generator/fake"
)

func TestInstructionFromAnswerRecipe(t *testing.T) {
	service := run.New(
		fake.New("What is Kubernetes?"),
	)

	output := filepath.Join(t.TempDir(), "dataset.jsonl")

	repoRoot, err := filepath.Abs("../../..")
	if err != nil {
		t.Fatal(err)
	}

	recipe := filepath.Join(
		repoRoot,
		"recipes",
		"instruction-from-answer.yaml",
	)

	err = service.Run(
		context.Background(),
		recipe,
		output,
	)

	if err != nil {
		t.Fatal(err)
	}
}
