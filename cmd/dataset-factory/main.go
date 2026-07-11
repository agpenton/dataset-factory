package main

import (
	"context"
	"log"

	"github.com/agpenton/dataset-factory/internal/bootstrap"
)

func main() {
	ctx := context.Background()

	application, err := app.New()
	if err != nil {
		log.Fatal(err)
	}

	if err := application.Run(ctx); err != nil {
		log.Fatal(err)
	}
}
