package dataset_test

import (
	"testing"

	"github.com/agpenton/dataset-factory/internal/domain/chunk"
	"github.com/agpenton/dataset-factory/internal/domain/dataset"
)

func TestFromChunks_CreatesJSONL(t *testing.T) {
	chunks, err := chunk.FromContent(
		"doc-1",
		"# Dataset Factory",
	)
	if err != nil {
		t.Fatal(err)
	}

	records, err := dataset.FromChunks(chunks)
	if err != nil {
		t.Fatal(err)
	}

	if len(records) != 1 {
		t.Fatalf("expected 1 record, got %d", len(records))
	}

	if records[0] == "" {
		t.Fatal("expected non-empty JSONL record")
	}
}
