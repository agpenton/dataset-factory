package recipe

import "errors"

var (
	ErrMissingAPIVersion = errors.New("missing apiVersion")
	ErrMissingKind       = errors.New("missing kind")
	ErrMissingName       = errors.New("missing metadata.name")
	ErrMissingVersion    = errors.New("missing metadata.version")
)

func Validate(r *Recipe) error {
	if r == nil {
		return errors.New("recipe is nil")
	}

	if r.APIVersion == "" {
		return ErrMissingAPIVersion
	}

	if r.Kind == "" {
		return ErrMissingKind
	}

	if r.Metadata.Name == "" {
		return ErrMissingName
	}

	if r.Metadata.Version == "" {
		return ErrMissingVersion
	}

	if r.Input.Type == "" {
		return errors.New("missing input.type")
	}

	if r.Input.Path == "" {
		return errors.New("missing input.path")
	}
	if r.Prompt.Template == "" {
		return errors.New("missing prompt.template")
	}

	return nil
}
