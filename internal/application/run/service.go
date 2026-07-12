package run

import (
	"context"

	"github.com/agpenton/dataset-factory/internal/application/export"
	apprecipe "github.com/agpenton/dataset-factory/internal/application/recipe"
)

type Service struct{}

func New() *Service {
	return &Service{}
}

func (s *Service) Run(
	ctx context.Context,
	recipePath string,
	outputPath string,
) error {
	_ = ctx

	r, err := apprecipe.Load(recipePath)
	if err != nil {
		return err
	}

	if err := apprecipe.Validate(r); err != nil {
		return err
	}

	// Temporary implementation.
	// This will be replaced by the recipe execution engine.
	records := []string{
		`{"messages":[]}`,
	}

	return export.JSONL(outputPath, records)
}
