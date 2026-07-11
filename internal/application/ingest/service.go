package ingest

import (
	"github.com/agpenton/dataset-factory/internal/domain/chunk"
	"github.com/agpenton/dataset-factory/internal/domain/dataset"
)

type Service struct{}

func New() *Service {
	return &Service{}
}

func (s *Service) Ingest(name, content string) ([]string, error) {
	chunks, err := chunk.FromContent(name, content)
	if err != nil {
		return nil, err
	}

	return dataset.FromChunks(chunks)
}
