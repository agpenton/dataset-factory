package recipe

type Recipe struct {
	APIVersion string    `yaml:"apiVersion"`
	Kind       string    `yaml:"kind"`
	Metadata   Metadata  `yaml:"metadata"`
	Input      Input     `yaml:"input"`
	Prompt     Prompt    `yaml:"prompt"`
	Pipeline   Pipeline  `yaml:"pipeline"`
	Output     Output    `yaml:"output"`
	Generator  Generator `yaml:"generator"`
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

type Pipeline struct {
	Type string `yaml:"type"`
}

type Output struct {
	Type string `yaml:"type"`
}

type Generator struct {
	Provider string `yaml:"provider"`
	Endpoint string `yaml:"endpoint"`
	Model    string `yaml:"model"`
}
