package bootstrap_test

import (
	"context"
	"testing"

	"github.com/agpenton/dataset-factory/internal/bootstrap"
)

func TestUnknownCommand(t *testing.T) {
	app, err := bootstrap.New()
	if err != nil {
		t.Fatal(err)
	}

	err = app.Run(
		context.Background(),
		[]string{"dataset-factory", "unknown"},
	)

	if err == nil {
		t.Fatal("expected error")
	}
}

func TestMissingRunArguments(t *testing.T) {
	app, err := bootstrap.New()
	if err != nil {
		t.Fatal(err)
	}

	err = app.Run(
		context.Background(),
		[]string{"dataset-factory", "run"},
	)

	if err == nil {
		t.Fatal("expected error")
	}
}
