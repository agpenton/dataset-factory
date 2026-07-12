package run

import (
	"context"
	"path/filepath"

	"github.com/agpenton/dataset-factory/internal/application/export"
	"github.com/agpenton/dataset-factory/internal/application/input"
	"github.com/agpenton/dataset-factory/internal/application/pipeline"
	apprecipe "github.com/agpenton/dataset-factory/internal/application/recipe"
	"github.com/agpenton/dataset-factory/internal/generator/fake"
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

	r, err := apprecipe.Load(recipePath)
	if err != nil {
		return err
	}

	if err := apprecipe.Validate(r); err != nil {
		return err
	}

	// TODO: Read this from the recipe once the schema supports it.
	recipeDir := filepath.Dir(recipePath)
	inputPath := filepath.Join(recipeDir, r.Input.Path)

	records, err := input.ReadAnswers(inputPath)
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}

	pipe := pipeline.NewInstructionFromAnswer(
		fake.New("What is Kubernetes?"),
	)

	datasetRecords, err := pipe.RunAll(ctx, records)
	if err != nil {
		return err
	}

	return export.JSONL(outputPath, datasetRecords)
}
