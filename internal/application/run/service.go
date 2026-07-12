package run

import (
	"context"

	apprecipe "github.com/agpenton/dataset-factory/internal/application/recipe"
)

type Service struct{}

func New() *Service {
	return &Service{}
}

func (s *Service) Run(ctx context.Context, recipePath string) error {
	_ = ctx

	r, err := apprecipe.Load(recipePath)
	if err != nil {
		return err
	}

	return apprecipe.Validate(r)
}
