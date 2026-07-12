package pipeline

import (
	"context"

	"github.com/agpenton/dataset-factory/internal/application/dataset"
	"github.com/agpenton/dataset-factory/internal/application/prompt"
	"github.com/agpenton/dataset-factory/internal/generator"
)

type InstructionFromAnswer struct {
	generator generator.Generator
}

func NewInstructionFromAnswer(g generator.Generator) *InstructionFromAnswer {
	return &InstructionFromAnswer{
		generator: g,
	}
}

func (p *InstructionFromAnswer) Run(
	ctx context.Context,
	answer string,
) (string, error) {

	promptText := prompt.BuildInstructionFromAnswer(answer)

	instruction, err := p.generator.Generate(ctx, promptText)
	if err != nil {
		return "", err
	}

	return dataset.BuildRecord(
		instruction,
		answer,
	), nil
}
