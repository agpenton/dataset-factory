package run_test

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/agpenton/dataset-factory/internal/application/run"
)

func TestInstructionFromAnswerRecipe(t *testing.T) {
	service := run.New()

	output := filepath.Join(t.TempDir(), "dataset.jsonl")

	err := service.Run(
		context.Background(),
		"../../../recipes/instruction-from-answer.yaml",
		output,
	)

	if err != nil {
		t.Fatal(err)
	}
}
