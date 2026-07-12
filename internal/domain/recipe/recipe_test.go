package recipe_test

import (
	"testing"

	"github.com/agpenton/dataset-factory/internal/domain/recipe"
)

func TestNewRecipe(t *testing.T) {
	r := recipe.New(
		"Instruction from Answer",
		"Generates instructions from seed answers.",
	)

	if r.Name() != "Instruction from Answer" {
		t.Fatal("unexpected recipe name")
	}

	if r.Description() != "Generates instructions from seed answers." {
		t.Fatal("unexpected description")
	}

	if r.Version() == "" {
		t.Fatal("expected version")
	}
}
