package filesystem

import (
	"os"
	"path/filepath"

	"github.com/agpenton/dataset-factory/internal/domain/document"
)

type Reader struct{}

func NewReader() *Reader {
	return &Reader{}
}

func (r *Reader) Read(path string) (*document.Document, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return document.New(
		filepath.Base(path),
		string(content),
	)
}
