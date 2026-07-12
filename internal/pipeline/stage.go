package pipeline

import "context"

type Stage[I any, O any] interface {
	Run(context.Context, I) (O, error)
}
