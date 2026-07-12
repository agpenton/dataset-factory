package recipe_test

import (
	"path/filepath"
	"testing"

	"github.com/agpenton/dataset-factory/internal/application/recipe"
)

func TestDebugRecipe(t *testing.T) {
	repoRoot, err := filepath.Abs("../../..")
	if err != nil {
		t.Fatal(err)
	}

	path := filepath.Join(
		repoRoot,
		"recipes",
		"instruction-from-answer.yaml",
	)

	t.Log(path)

	r, err := recipe.Load(path)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", r)
}
