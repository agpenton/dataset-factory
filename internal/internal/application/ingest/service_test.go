package ingest_test

import (
	"testing"

	"github.com/agpenton/dataset-factory/internal/application/ingest"
)

func TestService_Ingest(t *testing.T) {
	service := ingest.New()

	records, err := service.Ingest(
		"README.md",
		"# Dataset Factory",
	)
	if err != nil {
		t.Fatal(err)
	}

	if len(records) != 1 {
		t.Fatalf("expected 1 record, got %d", len(records))
	}

	if records[0] == "" {
		t.Fatal("expected non-empty dataset record")
	}
}
