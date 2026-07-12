package recipe

import "context"

type Recipe interface {
	Name() string
	Version() string

	Run(context.Context, Input) (Output, error)
}
