package recipe_test

import (
	"testing"

	"github.com/agpenton/dataset-factory/internal/application/recipe"
)

func TestLoad(t *testing.T) {
	r, err := recipe.Load("testdata/minimal.yaml")
	if err != nil {
		t.Fatal(err)
	}

	if r.APIVersion != "datasetfactory.io/v1alpha1" {
		t.Fatalf("unexpected apiVersion %q", r.APIVersion)
	}

	if r.Kind != "Recipe" {
		t.Fatalf("unexpected kind %q", r.Kind)
	}

	if r.Metadata.Name != "instruction-from-answer" {
		t.Fatalf("unexpected name %q", r.Metadata.Name)
	}

	if r.Metadata.Version != "v1" {
		t.Fatalf("unexpected version %q", r.Metadata.Version)
	}
}
