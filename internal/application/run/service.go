package run

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/agpenton/dataset-factory/internal/application/export"
	"github.com/agpenton/dataset-factory/internal/application/input"
	"github.com/agpenton/dataset-factory/internal/application/pipeline"
	apprecipe "github.com/agpenton/dataset-factory/internal/application/recipe"
	"github.com/agpenton/dataset-factory/internal/generator"
)

type Service struct {
	generator generator.Generator
}

func New(g generator.Generator) *Service {
	return &Service{
		generator: g,
	}
}
func (s *Service) Generator() generator.Generator {
	return s.generator
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

	fmt.Printf("recipePath = %q\n", recipePath)
	fmt.Printf("recipeDir  = %q\n", recipeDir)
	fmt.Printf("inputPath  = %q\n", inputPath)

	records, err := input.ReadAnswers(inputPath)
	if err != nil {
		return err
	}

	templatePath := filepath.Join(
		recipeDir,
		r.Prompt.Template,
	)

	pipe := pipeline.NewInstructionFromAnswer(
		s.generator,
		templatePath,
	)

	fmt.Printf("recipe: %s\n", recipePath)
	fmt.Printf("input : %s\n", inputPath)

	datasetRecords, err := pipe.RunAll(ctx, records)
	if err != nil {
		return err
	}

	return export.JSONL(outputPath, datasetRecords)
}
