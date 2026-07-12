package fake

import "context"

type Generator struct {
	response string
}

func New(response string) *Generator {
	return &Generator{
		response: response,
	}
}

func (g *Generator) Generate(
	ctx context.Context,
	prompt string,
) (string, error) {
	_ = ctx
	_ = prompt

	return g.response, nil
}
