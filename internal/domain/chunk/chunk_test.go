package chunk_test

import (
	"testing"

	"github.com/agpenton/dataset-factory/internal/domain/chunk"
)

func TestFromContent_CreatesSingleChunk(t *testing.T) {
	const (
		documentID = "document-1"
		content    = "# Dataset Factory"
	)

	chunks, err := chunk.FromContent(documentID, content)
	if err != nil {
		t.Fatal(err)
	}

	if len(chunks) != 1 {
		t.Fatalf("expected 1 chunk, got %d", len(chunks))
	}

	got := chunks[0]

	if got.DocumentID() != documentID {
		t.Fatalf("expected document ID %q, got %q",
			documentID,
			got.DocumentID(),
		)
	}

	if got.Content() != content {
		t.Fatalf("expected content %q, got %q",
			content,
			got.Content(),
		)
	}

	if got.ID() == "" {
		t.Fatal("expected non-empty chunk ID")
	}
}
