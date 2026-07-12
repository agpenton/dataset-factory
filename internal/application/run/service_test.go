package run_test

import (
	"context"
	"testing"

	"github.com/agpenton/dataset-factory/internal/application/run"
)

func TestService_RunRecipe(t *testing.T) {
	service := run.New()

	err := service.Run(
		context.Background(),
		"testdata/instruction-from-answer.yaml",
	)

	if err != nil {
		t.Fatal(err)
	}
}
