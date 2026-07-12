package run

import "context"

type Service struct{}

func New() *Service {
	return &Service{}
}

func (s *Service) Run(ctx context.Context, recipePath string) error {
	_ = ctx
	_ = recipePath

	return nil
}
