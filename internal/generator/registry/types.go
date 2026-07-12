package registry

import (
	apprecipe "github.com/agpenton/dataset-factory/internal/application/recipe"
	"github.com/agpenton/dataset-factory/internal/generator"
)

type Factory func(
	r *apprecipe.Recipe,
) (generator.Generator, error)
