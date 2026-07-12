package bootstrap

import (
	"context"
	"errors"
	"fmt"
)

type Application struct {
	container *Container
}

func New() *Application {
	return &Application{
		container: NewContainer(),
	}
}

func (a *Application) Run(ctx context.Context, args []string) error {
	_ = ctx

	if len(args) < 2 {
		return errors.New("missing command")
	}

	switch args[1] {
	case "run":
		if len(args) != 4 {
			return errors.New("usage: dataset-factory run <recipe.yaml> <output.jsonl>")
		}

		return a.container.RunService.Run(ctx, args[2], args[3])

	case "ingest":
		// TODO: implement
		return nil

	default:
		return fmt.Errorf("unknown command %q", args[1])
	}
}
