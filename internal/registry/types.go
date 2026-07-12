package registry

import "github.com/agpenton/dataset-factory/internal/pipeline"

type PipelineFactory func(
	ctx pipeline.Context,
) pipeline.Pipeline
