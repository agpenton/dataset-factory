package executor

import (
	"github.com/agpenton/dataset-factory/internal/application/input"
	"github.com/agpenton/dataset-factory/internal/domain/recipe"
	"github.com/agpenton/dataset-factory/internal/generator"
)

type Context struct {
	Recipe *recipe.Recipe

	Generator generator.Generator

	WorkingDir string

	Input []input.Record

	Current input.Record

	Prompt string

	Response string

	Dataset []string

	OutputPath string
}
