package pipeline

import (
	"github.com/agpenton/dataset-factory/internal/generator"
)

type Context struct {
	Generator  generator.Generator
	Template   string
	WorkingDir string
}
