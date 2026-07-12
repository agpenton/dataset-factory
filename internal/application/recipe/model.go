package recipe

type Recipe struct {
	APIVersion string   `yaml:"apiVersion"`
	Kind       string   `yaml:"kind"`
	Metadata   Metadata `yaml:"metadata"`
	Input      Input    `yaml:"input"`
	Prompt     Prompt   `yaml:"prompt"`
}

type Prompt struct {
	Template string `yaml:"template"`
}
type Metadata struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
}

type Input struct {
	Type string `yaml:"type"`
	Path string `yaml:"path"`
}
