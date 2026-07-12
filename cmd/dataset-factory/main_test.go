package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestRunRecipe(t *testing.T) {
	tmp := t.TempDir()

	output := filepath.Join(tmp, "dataset.jsonl")

	cmd := exec.Command(
		"go",
		"run",
		"./cmd/dataset-factory",
		"run",
		"./internal/application/recipe/testdata/minimal.yaml",
		output,
	)

	cmd.Env = os.Environ()

	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("%v\n%s", err, out)
	}

	if _, err := os.Stat(output); err != nil {
		t.Fatalf("expected %s to exist", output)
	}
}
