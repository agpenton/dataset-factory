package promptpreview

import (
	"context"
	"path/filepath"

	"github.com/agpenton/dataset-factory/internal/application/prompt"
	apprecipe "github.com/agpenton/dataset-factory/internal/application/recipe"
	"github.com/agpenton/dataset-factory/internal/generator"
)

type Service struct {
	renderer *prompt.Renderer
}

func (s *Service) Execute(
	ctx context.Context,
	recipePath string,
	answer string,
	g generator.Generator,
) (string, string, error) {

	rendered, err := s.Render(recipePath, answer)
	if err != nil {
		return "", "", err
	}

	response, err := g.Generate(ctx, rendered)
	if err != nil {
		return rendered, "", err
	}

	return rendered, response, nil
}

func New() *Service {
	return &Service{
		renderer: prompt.New(),
	}
}

func (s *Service) Render(
	recipePath string,
	answer string,
) (string, error) {

	r, err := apprecipe.Load(recipePath)
	if err != nil {
		return "", err
	}

	if err := apprecipe.Validate(r); err != nil {
		return "", err
	}

	recipeDir := filepath.Dir(recipePath)

	templatePath := filepath.Join(
		recipeDir,
		r.Prompt.Template,
	)

	return s.renderer.Render(
		templatePath,
		prompt.TemplateData{
			Answer: answer,
		},
	)
}
