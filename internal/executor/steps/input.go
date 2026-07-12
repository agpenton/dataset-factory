package steps

import (
	"context"

	"github.com/agpenton/dataset-factory/internal/application/input"
	"github.com/agpenton/dataset-factory/internal/executor"
)

type InputStep struct {
	Path string
}

func NewInput(path string) *InputStep {
	return &InputStep{
		Path: path,
	}
}

func (s *InputStep) Execute(
	_ context.Context,
	ctx *executor.Context,
) error {

	records, err := input.ReadAnswers(s.Path)
	if err != nil {
		return err
	}

	ctx.Input = records

	return nil
}
