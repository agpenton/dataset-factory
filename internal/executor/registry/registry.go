package registry

import (
	"fmt"

	"github.com/agpenton/dataset-factory/internal/executor"
)

var steps = map[string]Factory{}

func Register(name string, f Factory) {
	steps[name] = f
}

func New(name string, cfg map[string]any) (executor.Step, error) {
	f, ok := steps[name]
	if !ok {
		return nil, fmt.Errorf("unknown step %q", name)
	}

	return f(cfg)
}
