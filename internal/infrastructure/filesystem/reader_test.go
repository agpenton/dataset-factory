package filesystem_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/agpenton/dataset-factory/internal/infrastructure/filesystem"
)

func TestReader_ReadMarkdownFile(t *testing.T) {
	dir := t.TempDir()

	file := filepath.Join(dir, "README.md")

	err := os.WriteFile(
		file,
		[]byte("# Dataset Factory"),
		0o644,
	)
	if err != nil {
		t.Fatal(err)
	}

	reader := filesystem.NewReader()

	doc, err := reader.Read(file)
	if err != nil {
		t.Fatal(err)
	}

	if doc.Content() != "# Dataset Factory" {
		t.Fatal("unexpected content")
	}
}
