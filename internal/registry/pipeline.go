package registry

import (
	"fmt"
	"path/filepath"

	apprecipe "github.com/agpenton/dataset-factory/internal/application/recipe"
	"github.com/agpenton/dataset-factory/internal/generator"
	"github.com/agpenton/dataset-factory/internal/pipeline"
)

var pipelines = map[string]PipelineFactory{}

func Register(
	name string,
	factory PipelineFactory,
) {
	pipelines[name] = factory
}

func NewPipeline(
	r *apprecipe.Recipe,
	g generator.Generator,
	recipeDir string,
) (pipeline.Pipeline, error) {

	factory, ok := pipelines[r.Pipeline.Type]
	if !ok {
		return nil, fmt.Errorf(
			"unknown pipeline %q",
			r.Pipeline.Type,
		)
	}

	template := filepath.Join(
		recipeDir,
		r.Prompt.Template,
	)

	ctx := pipeline.Context{
		Generator:  g,
		Template:   template,
		WorkingDir: recipeDir,
	}

	return factory(ctx), nil
}
