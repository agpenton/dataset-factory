package recipe_test

import (
	"testing"

	"github.com/agpenton/dataset-factory/internal/application/recipe"
)

func TestValidateSuccess(t *testing.T) {
	r, err := recipe.Load("testdata/minimal.yaml")
	if err != nil {
		t.Fatal(err)
	}

	if err := recipe.Validate(r); err != nil {
		t.Fatal(err)
	}
}

func TestValidateMissingAPIVersion(t *testing.T) {
	r, _ := recipe.Load("testdata/missing-apiversion.yaml")

	if err := recipe.Validate(r); err == nil {
		t.Fatal("expected validation error")
	}
}

func TestValidateMissingKind(t *testing.T) {
	r, _ := recipe.Load("testdata/missing-kind.yaml")

	if err := recipe.Validate(r); err == nil {
		t.Fatal("expected validation error")
	}
}

func TestValidateMissingName(t *testing.T) {
	r, _ := recipe.Load("testdata/missing-name.yaml")

	if err := recipe.Validate(r); err == nil {
		t.Fatal("expected validation error")
	}
}

func TestValidateMissingVersion(t *testing.T) {
	r, _ := recipe.Load("testdata/missing-version.yaml")

	if err := recipe.Validate(r); err == nil {
		t.Fatal("expected validation error")
	}
}

func validRecipe() *recipe.Recipe {
	return &recipe.Recipe{
		APIVersion: "datasetfactory.io/v1alpha1",
		Kind:       "Recipe",
		Metadata: recipe.Metadata{
			Name:    "minimal",
			Version: "v1",
		},
		Pipeline: recipe.Pipeline{
			Type: "instruction-from-answer",
		},
		Input: recipe.Input{
			Type: "jsonl",
			Path: "test.jsonl",
		},
		Prompt: recipe.Prompt{
			Template: "prompt.tmpl",
		},
		Output: recipe.Output{
			Type: "chatml",
		},
	}
}
