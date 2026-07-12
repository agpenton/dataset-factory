package registry

import (
	apppipeline "github.com/agpenton/dataset-factory/internal/application/pipeline"
	"github.com/agpenton/dataset-factory/internal/pipeline"
)

func init() {
	Register(
		"instruction-from-answer",
		func(ctx pipeline.Context) pipeline.Pipeline {
			return apppipeline.NewInstructionFromAnswer(
				ctx.Generator,
				ctx.Template,
			)
		},
	)
}
