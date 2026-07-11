package document_test

import (
	"testing"

	"github.com/agpenton/dataset-factory/internal/domain/document"
)

func TestNewMetadata(t *testing.T) {
	m := document.NewMetadata(
		"docs/README.md",
		"text/markdown",
	)

	if got := m.Source(); got != "docs/README.md" {
		t.Fatalf("expected source %q, got %q", "docs/README.md", got)
	}

	if got := m.MediaType(); got != "text/markdown" {
		t.Fatalf("expected media type %q, got %q", "text/markdown", got)
	}
}
