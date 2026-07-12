package pipeline

import (
	"context"

	"github.com/agpenton/dataset-factory/internal/application/dataset"
	"github.com/agpenton/dataset-factory/internal/application/input"
	"github.com/agpenton/dataset-factory/internal/application/prompt"
	"github.com/agpenton/dataset-factory/internal/generator"
)

type InstructionFromAnswer struct {
	generator    generator.Generator
	templatePath string
	renderer     *prompt.Renderer
}

func NewInstructionFromAnswer(
	g generator.Generator,
	templatePath string,
) *InstructionFromAnswer {

	return &InstructionFromAnswer{
		generator:    g,
		templatePath: templatePath,
		renderer:     prompt.New(),
	}
}

func (p *InstructionFromAnswer) Run(
	ctx context.Context,
	answer string,
) (string, error) {

	promptText, err := p.renderer.Render(
		p.templatePath,
		prompt.TemplateData{
			Answer: answer,
		},
	)
	if err != nil {
		return "", err
	}
	if err != nil {
		return "", err
	}

	instruction, err := p.generator.Generate(ctx, promptText)
	if err != nil {
		return "", err
	}

	return dataset.BuildRecord(
		instruction,
		answer,
	), nil
}

func (p *InstructionFromAnswer) RunAll(
	ctx context.Context,
	records []input.Record,
) ([]string, error) {

	output := make([]string, 0, len(records))

	for _, record := range records {
		jsonl, err := p.Run(ctx, record.Answer)
		if err != nil {
			return nil, err
		}

		output = append(output, jsonl)
	}

	return output, nil
}
