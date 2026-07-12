package export_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/agpenton/dataset-factory/internal/application/export"
)

func TestExportJSONL(t *testing.T) {
	dir := t.TempDir()

	file := filepath.Join(dir, "dataset.jsonl")

	records := []string{
		`{"messages":[{"role":"user","content":"Hello"}]}`,
		`{"messages":[{"role":"assistant","content":"Hi"}]}`,
	}

	err := export.JSONL(file, records)
	if err != nil {
		t.Fatal(err)
	}

	data, err := os.ReadFile(file)
	if err != nil {
		t.Fatal(err)
	}

	expected :=
		`{"messages":[{"role":"user","content":"Hello"}]}
{"messages":[{"role":"assistant","content":"Hi"}]}
`

	if string(data) != expected {
		t.Fatalf("unexpected output\n%s", data)
	}
}
