package registry_test

import (
	"testing"

	apprecipe "github.com/agpenton/dataset-factory/internal/application/recipe"
	"github.com/agpenton/dataset-factory/internal/generator/fake"
	"github.com/agpenton/dataset-factory/internal/registry"
)

func TestUnknownPipeline(t *testing.T) {
	r := &apprecipe.Recipe{}

	r.Pipeline.Type = "does-not-exist"

	_, err := registry.NewPipeline(
		r,
		fake.New(""),
		t.TempDir(),
	)

	if err == nil {
		t.Fatal("expected error")
	}
}
