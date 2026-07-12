package main

import (
	"context"
	"log"
	"os"

	"github.com/agpenton/dataset-factory/internal/bootstrap"
)

func main() {
	ctx := context.Background()

	application := bootstrap.New()

	if err := application.Run(ctx, os.Args); err != nil {
		log.Fatal(err)
	}
}
