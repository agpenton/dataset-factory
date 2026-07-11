package app

import "context"

type App struct{}

func New() (*App, error) {
	return &App{}, nil
}

func (a *App) Run(ctx context.Context) error {
	return nil
}
