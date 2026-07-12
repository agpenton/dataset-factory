package registry

import (
	apprecipe "github.com/agpenton/dataset-factory/internal/application/recipe"
	"github.com/agpenton/dataset-factory/internal/generator"
	"github.com/agpenton/dataset-factory/internal/generator/fake"
)

func init() {
	Register(
		"fake",
		func(*apprecipe.Recipe) (generator.Generator, error) {
			return fake.New("What is Kubernetes?"), nil
		},
	)
}
