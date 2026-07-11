package document

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
)

var (
	ErrEmptyName    = errors.New("document name cannot be empty")
	ErrEmptyContent = errors.New("document content cannot be empty")
)

type Document struct {
	id      string
	name    string
	content string
}

func New(name, content string) (*Document, error) {
	if name == "" {
		return nil, ErrEmptyName
	}

	if content == "" {
		return nil, ErrEmptyContent
	}

	doc := &Document{
		name:    name,
		content: content,
	}

	doc.id = computeID(doc)

	return doc, nil
}

func (d *Document) ID() string {
	return d.id
}

func (d *Document) Name() string {
	return d.name
}

func (d *Document) Content() string {
	return d.content
}

func (d *Document) Validate() error {
	if d.name == "" {
		return ErrEmptyName
	}

	if d.content == "" {
		return ErrEmptyContent
	}

	return nil
}

func generateID(name, content string) string {
	hash := sha256.Sum256([]byte(name + "\n" + content))
	return hex.EncodeToString(hash[:])
}
