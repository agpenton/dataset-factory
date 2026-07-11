package chunk

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
)

package chunk

import (
"crypto/sha256"
"encoding/hex"
"errors"
)

var (
	ErrEmptyDocumentID = errors.New("document ID cannot be empty")
	ErrEmptyContent    = errors.New("content cannot be empty")
)

type Chunk struct {
	id         string
	documentID string
	content    string
}

func FromContent(documentID, content string) ([]Chunk, error) {
	if documentID == "" {
		return nil, ErrEmptyDocumentID
	}

	if content == "" {
		return nil, ErrEmptyContent
	}

	chunk := Chunk{
		id:         computeID(documentID, content),
		documentID: documentID,
		content:    content,
	}

	return []Chunk{chunk}, nil
}

func (c Chunk) ID() string {
	return c.id
}

func (c Chunk) DocumentID() string {
	return c.documentID
}

func (c Chunk) Content() string {
	return c.content
}

func computeID(documentID, content string) string {
	sum := sha256.Sum256([]byte(documentID + "\n" + content))
	return hex.EncodeToString(sum[:])
}