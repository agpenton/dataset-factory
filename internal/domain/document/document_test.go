package document_test

import (
	"testing"

	"github.com/agpenton/dataset-factory/internal/domain/document"
)

func TestNewDocument_CreatesValidDocument(t *testing.T) {
	doc, err := document.New(
		"README.md",
		"# Dataset Factory",
	)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if doc.ID() == "" {
		t.Fatal("expected document to have an ID")
	}

	if got := doc.Name(); got != "README.md" {
		t.Fatalf("expected name %q, got %q", "README.md", got)
	}

	if got := doc.Content(); got != "# Dataset Factory" {
		t.Fatalf("expected content %q, got %q", "# Dataset Factory", got)
	}

	if err := doc.Validate(); err != nil {
		t.Fatalf("expected document to be valid, got %v", err)
	}
}
func TestNewDocument_RejectsEmptyName(t *testing.T) {
	_, err := document.New("", "# Dataset Factory")

	if err == nil {
		t.Fatal("expected error")
	}
}

func TestNewDocument_RejectsEmptyContent(t *testing.T) {
	_, err := document.New("README.md", "")

	if err == nil {
		t.Fatal("expected error")
	}
}

func TestDocument_IDIsStable(t *testing.T) {
	doc1, err := document.New("README.md", "# Dataset Factory")
	if err != nil {
		t.Fatal(err)
	}

	doc2, err := document.New("README.md", "# Dataset Factory")
	if err != nil {
		t.Fatal(err)
	}

	if doc1.ID() != doc2.ID() {
		t.Fatal("expected deterministic ID")
	}
}

func TestDocument_IDIsDeterministic(t *testing.T) {
	doc1, err := document.New(
		"README.md",
		"# Hello",
	)
	if err != nil {
		t.Fatal(err)
	}

	doc2, err := document.New(
		"README.md",
		"# Hello",
	)
	if err != nil {
		t.Fatal(err)
	}

	if doc1.ID() != doc2.ID() {
		t.Fatalf(
			"expected IDs to match: %q != %q",
			doc1.ID(),
			doc2.ID(),
		)
	}
}

func TestDocument_DifferentDocumentsHaveDifferentIDs(t *testing.T) {
	doc1, err := document.New(
		"README.md",
		"# Hello",
	)
	if err != nil {
		t.Fatal(err)
	}

	doc2, err := document.New(
		"README2.md",
		"# Hello",
	)
	if err != nil {
		t.Fatal(err)
	}

	if doc1.ID() == doc2.ID() {
		t.Fatal("expected different IDs")
	}
}

func TestDocument_IDIsNeverEmpty(t *testing.T) {
	doc, err := document.New(
		"README.md",
		"# Hello",
	)
	if err != nil {
		t.Fatal(err)
	}

	if doc.ID() == "" {
		t.Fatal("expected non-empty ID")
	}
}
func TestDocument_WithMetadata(t *testing.T) {
	doc, err := document.New(
		"README.md",
		"# Dataset Factory",
	)
	if err != nil {
		t.Fatal(err)
	}

	meta := document.NewMetadata(
		"docs/README.md",
		"text/markdown",
	)

	updated := doc.WithMetadata(meta)

	if updated.Metadata() != meta {
		t.Fatal("expected metadata to be attached")
	}

	if doc.Metadata() == meta {
		t.Fatal("expected original document to remain unchanged")
	}
}

func TestDocument_Metadata(t *testing.T) {
	doc, err := document.New(
		"README.md",
		"# Dataset Factory",
	)
	if err != nil {
		t.Fatal(err)
	}

	meta := document.NewMetadata(
		"docs/README.md",
		"text/markdown",
	)

	doc = doc.WithMetadata(meta)

	got := doc.Metadata()

	if got.Source() != "docs/README.md" {
		t.Fatalf("expected source %q, got %q",
			"docs/README.md",
			got.Source(),
		)
	}

	if got.MediaType() != "text/markdown" {
		t.Fatalf("expected media type %q, got %q",
			"text/markdown",
			got.MediaType(),
		)
	}
}
