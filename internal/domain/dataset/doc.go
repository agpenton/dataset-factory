package dataset

import (
	"encoding/json"

	"github.com/agpenton/dataset-factory/internal/domain/chunk"
)

func FromChunks(chunks []chunk.Chunk) ([]string, error) {
	records := make([]string, 0, len(chunks))

	for _, c := range chunks {
		record := struct {
			ID         string `json:"id"`
			DocumentID string `json:"document_id"`
			Content    string `json:"content"`
		}{
			ID:         c.ID(),
			DocumentID: c.DocumentID(),
			Content:    c.Content(),
		}

		b, err := json.Marshal(record)
		if err != nil {
			return nil, err
		}

		records = append(records, string(b))
	}

	return records, nil
}
