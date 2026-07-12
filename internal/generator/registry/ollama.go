package registry

import (
	apprecipe "github.com/agpenton/dataset-factory/internal/application/recipe"
	"github.com/agpenton/dataset-factory/internal/generator"
	"github.com/agpenton/dataset-factory/internal/generator/ollama"
)

func init() {
	Register(
		"ollama",
		func(r *apprecipe.Recipe) (generator.Generator, error) {
			return ollama.New(
				r.Generator.Endpoint,
				r.Generator.Model,
			), nil
		},
	)
}
