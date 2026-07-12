package bootstrap

import (
	"github.com/agpenton/dataset-factory/internal/application/ingest"
	"github.com/agpenton/dataset-factory/internal/application/promptpreview"
	"github.com/agpenton/dataset-factory/internal/application/run"
	"github.com/agpenton/dataset-factory/internal/generator/ollama"
)

type Container struct {
	RunService    *run.Service
	IngestService *ingest.Service
	PromptService *promptpreview.Service
}

func NewContainer() *Container {
	return &Container{
		RunService: run.New(
			ollama.New(
				ollama.DefaultEndpoint,
				ollama.DefaultModel,
			),
		),
		IngestService: ingest.New(),
		PromptService: promptpreview.New(),
	}
}
