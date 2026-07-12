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

	case "prompt":
		if len(args) < 5 || args[3] != "--answer" {
			return errors.New(
				"usage: dataset-factory prompt <recipe.yaml> --answer <text> [--execute]",
			)
		}

		execute := len(args) == 6 && args[5] == "--execute"

		if !execute {
			rendered, err := a.container.PromptService.Render(
				args[2],
				args[4],
			)
			if err != nil {
				return err
			}

			fmt.Println(rendered)
			return nil
		}

		rendered, response, err := a.container.PromptService.Execute(
			ctx,
			args[2],
			args[4],
			a.container.RunService.Generator(),
		)
		if err != nil {
			return err
		}

		fmt.Println("========== Prompt ==========")
		fmt.Println(rendered)

		fmt.Println()
		fmt.Println("========== Response ==========")
		fmt.Println(response)

		return nil

	case "ingest":
		// TODO: implement
		return nil

	default:
		return fmt.Errorf("unknown command %q", args[1])
	}
}
