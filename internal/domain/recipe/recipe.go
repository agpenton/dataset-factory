package recipe

type Recipe struct {
	name        string
	description string
	version     string
}

func New(name, description string) Recipe {
	return Recipe{
		name:        name,
		description: description,
		version:     "v1",
	}
}

func (r Recipe) Name() string {
	return r.name
}

func (r Recipe) Description() string {
	return r.description
}

func (r Recipe) Version() string {
	return r.version
}
