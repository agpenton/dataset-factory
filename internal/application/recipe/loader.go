package recipe

import (
	"os"

	"gopkg.in/yaml.v3"
)

func Load(path string) (*Recipe, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var recipe Recipe

	if err := yaml.Unmarshal(data, &recipe); err != nil {
		return nil, err
	}

	return &recipe, nil
}
