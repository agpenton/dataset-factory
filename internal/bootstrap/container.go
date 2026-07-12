package bootstrap

import (
	"github.com/agpenton/dataset-factory/internal/application/ingest"
	"github.com/agpenton/dataset-factory/internal/application/run"
)

type Container struct {
	RunService    *run.Service
	IngestService *ingest.Service
}

func NewContainer() *Container {
	return &Container{
		RunService:    run.New(),
		IngestService: ingest.New(),
	}
}
