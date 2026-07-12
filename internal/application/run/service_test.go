package run_test

import (
	"context"
	"testing"

	"github.com/agpenton/dataset-factory/internal/application/run"
)

func TestRunRecipe(t *testing.T) {
	service := run.New()

	err := service.Run(
		context.Background(),
		"../recipe/testdata/minimal.yaml",
	)

	if err != nil {
		t.Fatal(err)
	}
}
