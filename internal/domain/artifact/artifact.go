package artifact

import "context"

type Artifact interface {
	ID() string
	Kind() Kind
	Metadata() Metadata
	Validate(context.Context) error
}
