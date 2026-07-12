package run_test

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/agpenton/dataset-factory/internal/application/run"
	"github.com/agpenton/dataset-factory/internal/generator/fake"
)

func TestRunRecipeProducesDataset(t *testing.T) {
	dir := t.TempDir()

	service := run.New(
		fake.New("What is Kubernetes?"),
	)

	output := filepath.Join(dir, "dataset.jsonl")

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

	if _, err := os.Stat(output); err != nil {
		t.Fatal("dataset was not generated")
	}
	data, err := os.ReadFile(output)
	if err != nil {
		t.Fatal(err)
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	if len(lines) != 2 {
		t.Fatalf("expected 2 dataset records, got %d", len(lines))
	}
}
