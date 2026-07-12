package run

import (
	"context"
	"path/filepath"

	"github.com/agpenton/dataset-factory/internal/application/export"
	"github.com/agpenton/dataset-factory/internal/application/input"
	apprecipe "github.com/agpenton/dataset-factory/internal/application/recipe"
	generatorregistry "github.com/agpenton/dataset-factory/internal/generator/registry"
	"github.com/agpenton/dataset-factory/internal/registry"
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

	recipeDir := filepath.Dir(recipePath)

	inputPath := filepath.Join(
		recipeDir,
		r.Input.Path,
	)

	records, err := input.ReadAnswers(inputPath)
	if err != nil {
		return err
	}

	g, err := generatorregistry.New(r)
	if err != nil {
		return err
	}

	pipe, err := registry.NewPipeline(
		r,
		g,
		recipeDir,
	)
	if err != nil {
		return err
	}

	datasetRecords, err := pipe.RunAll(ctx, records)
	if err != nil {
		return err
	}

	return export.JSONL(outputPath, datasetRecords)
}
