package run_test

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/agpenton/dataset-factory/internal/application/run"
)

func TestRunRecipeProducesDataset(t *testing.T) {
	dir := t.TempDir()

	service := run.New()

	output := filepath.Join(dir, "dataset.jsonl")

	err := service.Run(
		context.Background(),
		"../recipe/testdata/minimal.yaml",
		output,
	)
	if err != nil {
		t.Fatal(err)
	}

	if _, err := os.Stat(output); err != nil {
		t.Fatal("dataset was not generated")
	}
}
