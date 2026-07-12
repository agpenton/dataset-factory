package pipeline

import (
	"context"

	"github.com/agpenton/dataset-factory/internal/application/input"
)

type Pipeline interface {
	RunAll(
		ctx context.Context,
		records []input.Record,
	) ([]string, error)
}
