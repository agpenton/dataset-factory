package registry

import (
	"fmt"

	apprecipe "github.com/agpenton/dataset-factory/internal/application/recipe"
	"github.com/agpenton/dataset-factory/internal/generator"
)

var generators = map[string]Factory{}

func Register(
	name string,
	f Factory,
) {
	generators[name] = f
}

func New(
	r *apprecipe.Recipe,
) (generator.Generator, error) {

	factory, ok := generators[r.Generator.Provider]
	if !ok {
		return nil, fmt.Errorf(
			"unknown generator %q",
			r.Generator.Provider,
		)
	}

	return factory(r)
}
