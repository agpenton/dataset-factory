package executor_test

import (
	"context"
	"testing"

	"github.com/agpenton/dataset-factory/internal/executor"
)

type fakeStep struct {
	called bool
}

func (f *fakeStep) Execute(
	context.Context,
	*executor.Context,
) error {

	f.called = true
	return nil
}

func TestEngine(t *testing.T) {
	step := &fakeStep{}

	engine := executor.New(step)

	err := engine.Execute(
		context.Background(),
		&executor.Context{},
	)
	if err != nil {
		t.Fatal(err)
	}

	if !step.called {
		t.Fatal("step not executed")
	}
}
