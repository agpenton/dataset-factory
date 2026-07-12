package registry

import (
	"github.com/agpenton/dataset-factory/internal/executor"
)

type Factory func(map[string]any) (executor.Step, error)
